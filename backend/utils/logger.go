package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Logger provides application-wide logging functionality
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

// NewLogger creates a new Logger instance
func NewLogger(logFile string) (*Logger, error) {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(logFile), 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %v", err)
	}

	// Open log file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	// Create multi-writer for both file and stdout
	multiWriter := io.MultiWriter(os.Stdout, file)

	// Initialize loggers
	logger := &Logger{
		infoLogger:  log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime),
		errorLogger: log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(multiWriter, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	return logger, nil
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	l.debugLogger.Printf(format, v...)
}

// LogRequest logs an HTTP request
func (l *Logger) LogRequest(method, path, ip string, duration time.Duration) {
	l.Info("Request: %s %s from %s took %v", method, path, ip, duration)
}

// LogError logs an error with context
func (l *Logger) LogError(err error, context string) {
	l.Error("%s: %v", context, err)
}

// LogMarginCall logs a margin call alert
func (l *Logger) LogMarginCall(clientID int64, portfolioValue, netEquity, marginShortfall float64) {
	l.Error("MARGIN CALL ALERT - Client ID: %d", clientID)
	l.Error("Portfolio Value: $%.2f", portfolioValue)
	l.Error("Net Equity: $%.2f", netEquity)
	l.Error("Margin Shortfall: $%.2f", marginShortfall)
	l.Error("Time: %s", time.Now().Format(time.RFC3339))
}

// Close closes the log file
func (l *Logger) Close() error {
	// The log.Logger doesn't need to be closed as it uses the underlying writer
	// which is managed by the caller
	return nil
}
