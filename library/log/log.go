package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Config struct {
	DEV        bool
	FILE       string
	MAXSIZE    int
	MAXBACKUPS int
	MAXAGE     int
}

var (
	Log *zap.SugaredLogger
)

func New(c *Config) {
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
	// add caller and remove wrapped file
	options = append(options, zap.AddCallerSkip(1), zap.AddCaller())
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
	Log = zap.New(coreTee, options...).Sugar()
}

func Info(args ...interface{}) {
	Log.Info(args)
}

func Infof(template string, args ...interface{}) {
	Log.Infof(template, args)
}

func Fatalf(template string, args ...interface{}) {
	Log.Fatalf(template, args)
}
