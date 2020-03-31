package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Config struct {
	DEV bool
	// Filebeat input path
	FILE       string
	MAXSIZE    int
	MAXBACKUPS int
	MAXAGE     int
}

var (
	Logger *zap.SugaredLogger
)

func NewLogger(c *Config) {
	var options []zap.Option

	fileHandler := zapcore.AddSync(&lumberjack.Logger{
		Filename:   c.FILE,
		MaxSize:    c.MAXSIZE,
		MaxBackups: c.MAXBACKUPS,
		MaxAge:     c.MAXAGE,
	})

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// add caller
	options = append(options, zap.AddCaller())
	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	coreTee := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			fileHandler,
			zap.DebugLevel,
		),
		zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.Lock(os.Stdout), zap.DebugLevel),
	)
	Logger = zap.New(coreTee, options...).Sugar()
}
