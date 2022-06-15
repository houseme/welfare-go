package log

import (
	"context"
	"errors"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/houseme/welfare-go/pkg"
)

var (
	ErrInvalidKey = errors.New("invalid key")
)

type Level zapcore.Level

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = Level(zap.DebugLevel)
	// InfoLevel is the default logging priority.
	InfoLevel = Level(zap.InfoLevel)
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = Level(zap.WarnLevel)
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = Level(zap.ErrorLevel)
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = Level(zap.DPanicLevel)
	// PanicLevel logs a message, then panics.
	PanicLevel = Level(zap.PanicLevel)
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = Level(zapcore.FatalLevel)
)

// Logger is the global logger instance.
type Logger struct {
	Level Level `json:"level"`
	Log   *zap.Logger
}

// New is the global logger instance.
func New(ctx context.Context, logPath string, loglevel Level) *Logger {
	var coreArr []zapcore.Core

	if pkg.Helper().Trim(logPath) != "" {
		logPath = os.TempDir()
	}

	// 获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()            // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder       // 显示完整文件路径
	encoderConfig.EncodeDuration = zapcore.MillisDurationEncoder // 指定时间格式
	encoderConfig.EncodeName = zapcore.FullNameEncoder           // 显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)          // 创建一个文件输出器，参数是指定文件路径，不存在则创建

	// 日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // error级别
		return lev >= zap.ErrorLevel
	})
	if loglevel <= InfoLevel {
		lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // info和debug级别,debug级别是最低的
			return lev < zap.ErrorLevel && lev >= zap.DebugLevel
		})

		// info文件writeSyncer
		infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath + "/log/info.log", // 日志文件存放目录，如果文件夹不存在会自动创建
			MaxSize:    2,                         // 文件大小限制,单位MB
			MaxBackups: 100,                       // 最大保留日志文件数量
			MaxAge:     30,                        // 日志文件保留天数
			Compress:   false,                     // 是否压缩处理
		})
		infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
		coreArr = append(coreArr, infoFileCore)
	}
	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath + "/log/error.log", // 日志文件存放目录
		MaxSize:    1,                          // 文件大小限制,单位MB
		MaxBackups: 5,                          // 最大保留日志文件数量
		MaxAge:     30,                         // 日志文件保留天数
		Compress:   false,                      // 是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	coreArr = append(coreArr, errorFileCore)
	return &Logger{
		Level: loglevel,
		Log:   zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()), // zap.AddCaller()为显示文件名和行号，可省略
	}
}
