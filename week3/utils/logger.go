package utils

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(logPath string) {
	// Lumberjack is a log rotation library
	// Fluentbit is a log aggregator (just takes logs from multiple places and sends somewhere)
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    20,
		MaxBackups: 4, // number of old log files to keep
		MaxAge:     30,
		Compress:   false, // or Fluentbit might not pick up
	}

	// Step 2.
	// Define where to synchronize each file, since these are "Syncers".
	// And we want a multi-writer.
	fileSync := zapcore.AddSync(lumberjackLogger)
	stdoutSync := zapcore.AddSync(os.Stdout)
	coreSync := zapcore.NewMultiWriteSyncer(fileSync, stdoutSync)

	// Step 3.
	// Define the config.
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.TimeKey = "timestamp"
	encoder := zapcore.NewJSONEncoder(config)

	core := zapcore.NewCore(encoder, coreSync, zapcore.DebugLevel)

	// Step 4.
	// Initialize the logger.
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	defer logger.Sync()
}

func GetLogger() *zap.Logger {
	return logger
}
