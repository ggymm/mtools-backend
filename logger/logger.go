package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(appPath string) (*zap.SugaredLogger, func(), error) {
	db, cleanFunc, err := NewLogger(appPath)
	if err != nil {
		return nil, cleanFunc, err
	}
	return db, cleanFunc, nil
}

func NewLogger(appPath string) (*zap.SugaredLogger, func(), error) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: appPath + "log/mtools-backend.log",
		MaxAge:   30,
		MaxSize:  256,
		Compress: true,
	})
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		atomicLevel,
	)
	Sugar := zap.New(core).Sugar()
	cleanFunc := func() {
		_ = Sugar.Sync()
	}
	return Sugar, cleanFunc, nil
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
