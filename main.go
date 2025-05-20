package main

import (
	"prirate-lang-go/core/logger"
	"prirate-lang-go/core/server"
)

func main() {
	// Echo instance

	if err := server.Run(); err != nil {
		logger.Error("run server error", err)
	}
}
