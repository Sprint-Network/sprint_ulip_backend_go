package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(logFilePath string) {
	Log = logrus.New()

	// Set output to both a file and console
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		Log.Fatalf("Failed to open log file: %s", err)
	}

	// Multi-output to file and console
	Log.SetOutput(file) // Comment this line if multi-output is not needed

	// Log Format
	Log.SetFormatter(&logrus.JSONFormatter{ // For structured JSON logs
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log.SetFormatter(&logrus.TextFormatter{ // For human-readable text logs
		FullTimestamp: true,
	})

	// Set log level (Adjust as needed: DebugLevel, InfoLevel, WarnLevel, ErrorLevel)
	Log.SetLevel(logrus.DebugLevel)

	// Example message after initialization
	Log.Info("Logger initialized")
}
