package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// newZapLogger return a zap logger.
func newZap(encoder zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *zap.Logger {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
		), level)
	return zap.New(core, opts...)
}

func NewZapLogger() *zap.Logger {
	encoder := zapcore.EncoderConfig{
		TimeKey:        "t",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	logger := newZap(
		encoder,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.Development(),
	)
	return logger
}

//// Log Implementation of logger interface.
//func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
//	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
//		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
//		return nil
//	}
//	// Zap.Field is used when keyvals pairs appear
//	var data []zap.Field
//	for i := 0; i < len(keyvals); i += 2 {
//		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
//	}
//	switch level {
//	case log.LevelDebug:
//		l.log.Debug("", data...)
//	case log.LevelInfo:
//		l.log.Info("", data...)
//	case log.LevelWarn:
//		l.log.Warn("", data...)
//	case log.LevelError:
//		l.log.Error("", data...)
//	}
//	return nil
//}
