package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB              string        `mapstructure:"DB"`
	Port            string        `mapstructure:"PORT"`
	DBName          string        `mapstructure:"DB_NAME"`
	TokenPrivateKey string        `mapstructure:"TOKEN_PRIVATE_KEY"`
	TokenPublicKey  string        `mapstructure:"TOKEN_PUBLIC_KEY"`
	TokenExpiration time.Duration `mapstructure:"TOKEN_EXPIRATION"`
	TokenMaxAge     int           `mapstructure:"TOKEN_MAXAGE"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)

	return config, err
}
