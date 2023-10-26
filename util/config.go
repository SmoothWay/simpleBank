package util

import "github.com/spf13/viper"

type Config struct {
	DBDRiver      string `mapstructure:"DB_DRIVER"`
	DSN           string `mapstructure:"DB_DSN"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
