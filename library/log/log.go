package log

import (
	"fmt"
	"go.uber.org/zap"
)

type Config struct {
	DEV bool
}

var (
	Log *zap.SugaredLogger
)

func New(c *Config) {
	var (
		logger         *zap.Logger
		err            error
		newConstructor func(options ...zap.Option) (*zap.Logger, error)
		options        []zap.Option
	)
	options = append(options, zap.AddCallerSkip(1))
	if c.DEV {
		newConstructor = zap.NewDevelopment
	} else {
		newConstructor = zap.NewProduction
	}
	if logger, err = newConstructor(options...); err != nil {
		panic(fmt.Sprintf("logger init failed %s", err.Error()))
	}
	Log = logger.Sugar()
	defer Log.Sync()
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
