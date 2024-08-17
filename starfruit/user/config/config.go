package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"os"
)

var AppConfig = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
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
	cfg.viper.AddConfigPath(getWd() + "/user/config")

	// 读取文件
	err := cfg.viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	cfg.ReadServerConfig()
	return cfg
}

// ReadServerConfig 读取配置文件
func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")

	c.SC = sc
}

// ReadRedisConfig 读取redis配置文件
func (c *Config) ReadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("redis.db"),
	}
}
