package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewConfig 初始化viper設定檔
func SetConfig() {
	viper.SetConfigType("toml")
	viper.AutomaticEnv()

	// read project default config first
	viper.AddConfigPath("./configs")
	viper.SetConfigName("dispatcher")

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(fmt.Errorf("Fatal error default config file: %s", err))
	}
}
