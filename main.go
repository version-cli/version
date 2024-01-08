/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package main

import (
	"needs-a-name/cmd"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	if os.Getenv("LOG_LEVEL") == "debug" {
		config = zap.NewDevelopmentConfig()
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zap.Must(config.Build())
}

func init() {
	logger := createLogger()

	defer logger.Sync()

	zap.ReplaceGlobals(logger)

}

func main() {
	cmd.Execute()
}
