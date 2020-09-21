package logger

import (
	l "log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	config := zap.Config{
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:         "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = config.Build(); err != nil {
		l.Fatalln(err)
	}
}

func GetLogger() *zap.Logger {
	return log
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	_ = log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.Error(err))
	log.Error(msg, tags...)
	_ = log.Sync()
}
