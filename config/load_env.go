package config

import "github.com/spf13/viper"

type Config struct {
	DBhost string `mapstructure:"SQL_HOST"`
	DBuser string `mapstructure:"SQL_USER"`
	DBpass string `mapstructure:"SQL_PASS"`
	DBname string `mapstructure:"SQL_DB"`
	DBport string `mapstructure:"SQL_PORT"`

	RedisURL string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
