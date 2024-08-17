package config

import (
	"github.com/spf13/viper"
	"os"
)

var AppConfig = InitConfig()

type Config struct {
	viper *viper.Viper
}

// ServerConfig 服务配置
type ServerConfig struct {
	Name string
	Addr string
}

// getWd 获取工作目录
func getWd() string {
	dir, _ := os.Getwd()
	return dir
}

// InitConfig 初始化配置文件
func InitConfig() *Config {
	cfg := &Config{viper: viper.New()}
	cfg.viper.SetConfigName("config")
	cfg.viper.SetConfigType("yaml")
	cfg.viper.AddConfigPath(getWd() + "/config")

	// 读取文件
	err := cfg.viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
