// Package log 提供一個uber zap library的簡單封裝
package log

import (
	"go.uber.org/zap"
	"github.com/spf13/viper"
)

var l *zap.Logger

// InitLog init zap log
func InitLog() (*zap.Logger, error) {
	var cfg zap.Config

	if disableJSON := viper.GetBool("log.disable_json"); disableJSON {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	cfg.OutputPaths = []string{
		viper.GetString("log.file"),
	}

	return cfg.Build()
}

// NewLog init zap log
// 如果debug打開 輸出console好閱讀的模式, 輸出Debug級以上的log
// 如果debug關閉, 輸出Info級以上的log, 且格式化成json, 方便收集到ELK上
func NewLog(outputPaths []string, debug bool) (*zap.Logger, error) {
	var cfg zap.Config
	var err error

	if debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	cfg.OutputPaths = outputPaths

	l, err = cfg.Build()
	return l, err
}

// GetLog gets the global zap logger instance.
func GetLog() *zap.Logger {
	return l
}
