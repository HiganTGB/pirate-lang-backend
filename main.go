package main

import (
	"pirate-lang-go/core/logger"
	"pirate-lang-go/core/server"
)

func main() {
	// Echo instance

	if err := server.Run(); err != nil {
		logger.Error("run server error", err)
	}
}
