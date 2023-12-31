package utils

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 11:54
//

import (
	"backend/app/config"
	"backend/pkg/utils/files"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLogger(zapCfg config.Zap) *zap.Logger {
	// 生成日志文件目录
	if ok, _ := files.PathExists(zapCfg.Directory); !ok {
		log.Printf("create %v directory\n", zapCfg.Directory)
		_ = os.Mkdir(zapCfg.Directory, os.ModePerm)
	}
	core := zapcore.NewCore(getEncoder(zapCfg.Format), getWriterSyncer(zapCfg), getLevelPriority(zapCfg))
	Logger = zap.New(core)

	if zapCfg.ShowLine {
		// 获取 调用的文件, 函数名称, 行号
		Logger = Logger.WithOptions(zap.AddCaller())
	}

	log.Println("Zap Logger 初始化成功")
	return Logger
}

// 编码器: 如何写入日志
func getEncoder(format string) zapcore.Encoder {
	// 参考: zap.NewProductionEncoderConfig()
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // ?
	}

	if format == "json" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志输出路径: 文件、控制台、双向输出s
func getWriterSyncer(zapCfg config.Zap) zapcore.WriteSyncer {
	file, _ := os.Create(zapCfg.Directory + "/log.log")

	// 双向输出
	if zapCfg.LogInConsole {
		fileWriter := zapcore.AddSync(file)
		consoleWriter := zapcore.AddSync(os.Stdout)
		return zapcore.NewMultiWriteSyncer(fileWriter, consoleWriter)
	}

	// 输出到文件
	return zapcore.AddSync(file)
}

// 获取日志输出级别
func getLevelPriority(zapCfg config.Zap) zapcore.LevelEnabler {
	switch zapCfg.Level {
	case "debug", "Debug":
		return zap.DebugLevel
	case "info", "Info":
		return zap.InfoLevel
	case "warn", "Warn":
		return zap.WarnLevel
	case "error", "Error":
		return zap.ErrorLevel
	case "dpanic", "DPanic":
		return zap.DPanicLevel
	case "panic", "Panic":
		return zap.PanicLevel
	case "fatal", "Fatal":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006/01/02 - 15:04:05"))
}
