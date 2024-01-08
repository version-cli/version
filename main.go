/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"needs-a-name/cmd"
	"os"
)

func createLogger() *zap.Logger {
	//return zap.Must(config.Build())
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
