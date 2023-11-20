package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresURL string `mapstructure:"POSTGRES_URL"`
	APIPort     string `mapstructure:"API_PORT"`
	FileConfig
	TokenConfig
}

type FileConfig struct{}

type TokenConfig struct{}

func NewConfig() Config {
	cfg := Config{}
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("☠️ Environment can't be loaded: ", err)
	}
	return cfg
}
