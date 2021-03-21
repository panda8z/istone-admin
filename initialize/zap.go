package initialize

import (
	"fmt"
	"github.com/panda8z/istone/global"
	"github.com/panda8z/istone/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.ISA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n",
			global.ISA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.ISA_CONFIG.Zap.Director, os.ModePerm)
	}

	switch global.ISA_CONFIG.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	case "info":
		level = zap.InfoLevel
	case "panic":
		level = zap.PanicLevel
	case "warn":
		level = zap.WarnLevel
	default:
		level = zap.InfoLevel
	}
	var z *zap.Logger
	if level == zap.InfoLevel || level == zap.ErrorLevel {
		z = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.ISA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return z
}

func encodeConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    "function",
		StacktraceKey:  global.ISA_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.ISA_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.ISA_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.ISA_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.ISA_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.ISA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(encodeConfig())
	}
	return zapcore.NewConsoleEncoder(encodeConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := utils.GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.ISA_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
