package database

import (
	"database/sql"
	"fmt"
	"prirate-lang-go/core/logger"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type IDatabase interface {
	DB() *sql.DB
}

type Database struct {
	db *sql.DB
}

type DatabaseConfig struct {
	Host                   string
	Port                   int
	User                   string
	Password               string
	DBName                 string
	MaxOpenConns           int
	MaxIdleConns           int
	ConnMaxLifetime        int    // in minutes
	SSLMode                string // disable, require, verify-ca, verify-full
	ConnectTimeout         int    // in seconds
	StatementTimeout       int    // in seconds
	IdleInTxSessionTimeout int    // in seconds
}

var (
	instance *Database
	once     sync.Once
)

func GetDB() IDatabase {
	return instance
}

func InitDB(config DatabaseConfig) (Database, error) {
	logger.Info("Initializing database...")
	var db Database
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		return Database{}, fmt.Errorf("failed to connect to database: %w", err)
	}
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute)

	if err = sqlDB.Ping(); err != nil {
		logger.Error("Failed to ping database", "error", err)
		return Database{}, fmt.Errorf("failed to ping database: %w", err)
	}

	db = Database{
		db: sqlDB,
	}

	logger.Info("Database initialized successfully",
		"maxOpenConns", config.MaxOpenConns,
		"maxIdleConns", config.MaxIdleConns,
		"connMaxLifetime", config.ConnMaxLifetime,
	)

	return db, nil
}
func (d *Database) DB() *sql.DB {
	return d.db
}
