package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config 配置
type Config struct {
	*viper.Viper
}

// New 创建配置
func New() *Config {
	viper := viper.New()
	// 优先使用环境变量
	viper.AutomaticEnv()
	// database.mysql.dns => DATABASE_MYSQL_DNS
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	return &Config{viper}
}
