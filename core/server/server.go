package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin" // Replace Echo with Gin
	"prirate-lang-go/core/config"
	"prirate-lang-go/core/database"
	"prirate-lang-go/core/logger"
	"prirate-lang-go/core/middleware" // Middleware already converted to Gin
	"prirate-lang-go/modules/account"
)

type Server struct {
	ginEngine *gin.Engine // Replace echo.Echo with *gin.Engine
	addr      string
	db        database.Database
}

func initEnvironment() (config.Environment, error) {
	env := flag.String("env", "dev", "Environment (dev/prod)")
	flag.Parse()

	switch *env {
	case "dev":
		return config.DevEnvironment, nil
	case "prod":
		return config.ProdEnvironment, nil
	default:
		return "", fmt.Errorf("invalid environment. Use 'dev' or 'prod'")
	}
}

func initServer() (*Server, error) {
	environment, err := initEnvironment()
	if err != nil {
		return nil, err
	}

	if errInitConfig := config.Init(environment); errInitConfig != nil {
		return nil, fmt.Errorf("failed to initialize config: %w", errInitConfig)
	}

	cfg := config.Get()

	// Initialize logger
	if errInitLogger := logger.Init(logger.LogConfig{
		Level:    logger.LogLevelDebug,
		FilePath: "app.log",
	}); errInitLogger != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", errInitLogger)
	}

	// Initialize database
	db, err := database.InitDB(database.DatabaseConfig{
		Host:                   cfg.Database.Host,
		Port:                   cfg.Database.Port,
		User:                   cfg.Database.User,
		Password:               cfg.Database.Password,
		DBName:                 cfg.Database.DBName,
		MaxOpenConns:           10, // Default value
		MaxIdleConns:           5,  // Default value
		ConnMaxLifetime:        60, // Default value in minutes
		SSLMode:                "disable",
		ConnectTimeout:         10,
		StatementTimeout:       30,
		IdleInTxSessionTimeout: 60,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	logger.Info("Server initializing",
		"environment", environment,
		"host", cfg.Server.Host,
		"port", cfg.Server.Port,
	)

	// Initialize Gin Engine
	// gin.Default() includes the default Logger and Recovery middleware.
	// If you want full control over the middleware, use gin.New()
	// and add gin.Recovery() yourself to handle panics.
	r := gin.New()

	// Middleware
	r.Use(middleware.LoggerMiddleware())

	// Initialize modules
	// The account.Init function needs to be adjusted to accept *gin.Engine instead of *echo.Echo
	account.Init(r, db)

	return &Server{
		ginEngine: r, // Assign gin.Engine to the Server struct
		addr:      fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		db:        db,
	}, nil
}

func (s *Server) start() error {
	logger.Info("Starting HTTP server", "address", s.addr)

	// Create an http.Server to manage server shutdown
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.ginEngine, // Assign the Gin Engine to http.Server
	}

	go func() {
		// Start listening and serving requests
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Log the error if the server fails to listen or the error is not due to closing the server
			logger.Error("Could not listen on", "address", s.addr, "error", err)
		}
	}()

	// Wait for an interrupt signal (Ctrl+C or kill command)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server gracefully: %w", err)
	}

	logger.Info("Server shutdown complete")
	return nil
}

func Run() error {
	srv, err := initServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err) // Use log.Fatalf to exit if initServer fails
		return err                                         // This line might not be necessary after using Fatalf
	}
	return srv.start()
}
