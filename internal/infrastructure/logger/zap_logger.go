package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type ZapLogger struct {
	logger *zap.Logger
}

func localTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.In(time.Local).Format("2006-01-02T15:04:05.000Z07:00"))
}

func NewProductionLogger(verbose bool, name string) (*zap.Logger, error) {
	logFile := fmt.Sprintf("logs/%s.log", name)
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     localTimeEncoder, // Use local timezone
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logLevel := zapcore.InfoLevel
	if verbose {
		logLevel = zapcore.DebugLevel
	}

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(file),
		logLevel,
	)

	stdoutCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)
	core := zapcore.NewTee(fileCore, stdoutCore)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger, nil
}

func (zl *ZapLogger) Info(msg string, fields ...any) {
	zl.logger.Sugar().Infow(msg, fields)
}

func (zl *ZapLogger) Error(msg string, fields ...any) {
	zl.logger.Sugar().Errorw(msg, fields)
}
