package public

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出方式
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名

	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日志留存时间
	ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}


// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch z.EncodeLevel  {
	case "LowercaseLevelEncoder": // 小写编译器(默认)
		return zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": //带颜色小写编译器
		return zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // 大写编译器
		return zapcore.CapitalLevelEncoder
	case "CapitalColorEncoder": // 带颜色大写编译器
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}


// 根据字符串转换为 zapcore.Level
func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
