package initialize

import (
	"bullet-screen/common"
	"bullet-screen/model/public"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var Z = new(_zap)
var z public.Zap

type _zap struct {}

func ZapInit() *zap.Logger {
	z = public.Zap{
		Level: common.Global_Viper.GetString("zap.level"),
		Prefix: common.Global_Viper.GetString("zap.prefix"),
		Format: common.Global_Viper.GetString("zap.format"),
		Director: common.Global_Viper.GetString("zap.director"),
		EncodeLevel: common.Global_Viper.GetString("zap.encode-level"),
		StacktraceKey: common.Global_Viper.GetString("zap.stacktrace-key"),
		MaxAge: common.Global_Viper.GetInt("zap.max-age"),
		ShowLine: common.Global_Viper.GetBool("zap.show-line"),
		LogInConsole: common.Global_Viper.GetBool("zap.log-in-console"),
	}

	cores := Z.GetZapCores()
	logger := zap.New(zapcore.NewTee(cores...))

	if z.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// 获取zapcore.Encoder
func (Z *_zap) GetEncoder() zapcore.Encoder {
	if z.Format == "json" {
		return zapcore.NewJSONEncoder(Z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(Z.GetEncoderConfig())
}

// 获取zapcore.EncoderConfig
func (Z *_zap)  GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey: "message",
		//LevelKey: "level",
		TimeKey: "time",
		NameKey: "logger",
		CallerKey: "caller",
		StacktraceKey: z.StacktraceKey,
		LineEnding: zapcore.DefaultLineEnding,
		//EncodeLevel: z.ZapEncodeLevel,
		EncodeTime: Z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}
}

// CustomTimeEncoder 自定义日志输出时间格式
// Author [SliverHorn](https://github.com/SliverHorn)
func (Z *_zap)  CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(z.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

func (Z *_zap)  GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(z.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(z.MaxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if z.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

// 获取Encoder的zapcore.Core
func (Z *_zap)  GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := Z.GetWriteSyncer(l.String())
	if err != nil {
		fmt.Println("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(Z.GetEncoder(), writer, level)
}


// GetZapCores 根据配置文件的Level获取 []zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (Z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, Z.GetEncoderCore(level, Z.GetLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
// Author [SliverHorn](https://github.com/SliverHorn)
func (Z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}
