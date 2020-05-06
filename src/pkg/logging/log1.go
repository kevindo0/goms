package log

import (
	"os"
	"txcb/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func Setup(cfg *config.Config) {
	Debug := true
	l := lumberjack.Logger{
		Filename:   cfg.File,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   false,
	}

	// 设置日志级别
	var level zap.AtomicLevel
	var writeSyncer zapcore.WriteSyncer
	var encoder zapcore.Encoder

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	if Debug {
		// debug模式
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		// 开启日志颜色模式，并以console形式输出
		encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		// 打印到控制台和文件
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&l))
	} else {
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		// 以json模式输出日志
		encoder = zapcore.NewJSONEncoder(encoderConfig)
		// 打印到文件
		writeSyncer = zapcore.AddSync(&l)
	}

	core := zapcore.NewCore(
		encoder, // 编码器配置
		writeSyncer,
		level,
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	skip := zap.AddCallerSkip(1)
	levelError := zapcore.ErrorLevel
	trace := zap.AddStacktrace(levelError)
	error_out := zap.ErrorOutput(writeSyncer)

	logger = zap.New(core, caller, skip, trace, error_out)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Sync() {
	logger.Sync()
}

func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func ZError(err error) zap.Field {
	return zap.Error(err)
}

func String(key string, val string) zap.Field {
	return zap.String(key, val)
}

func Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}
