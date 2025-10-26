// package main
package main

import (
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	if err := Run(); err != nil {
		logger.Error("Error runnig server", zap.Any("error", err))
		os.Exit(1)
	}
	logger.Info("All services are offline")
}
