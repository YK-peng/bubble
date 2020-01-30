package xlog

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//全局业务层log
var logger Xlogger

func NewGbizLogger(logPath string, logLevel string) {
	logger = NewLogger(logPath, logLevel)
}

//@logpath 日志文件路径
//@loglevel 日志级别
func NewLogger(logPath string, logLevel string) *zap.SugaredLogger {
	hook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    128,     // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 30,      // 保留旧文件的最大个数
		MaxAge:     7,       // 保留旧文件的最大天数
		Compress:   true,    // 是否压缩/归档旧文件 disabled by default
	}

	w := zapcore.AddSync(&hook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		// 有好的格式、输出控制台、动态等级
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), os.Stdout, level),
		// json格式、输出文件、处定义等级规则
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), w, level),
	)

	logger := zap.New(core)
	sugarLogger := logger.Sugar()
	sugarLogger.Info("DefaultLogger init success")

	//服务重新启动，日志会追加，不会删除
	return sugarLogger
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

//func Warning(args ...interface{}) {
//	logger.Warn(args...)
//}
//
//func Warningf(format string, args ...interface{}) {
//	logger.Warnf(format, args...)
//}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
	os.Exit(1)
}
