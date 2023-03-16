package configs

import "github.com/spf13/viper"

type Config struct {
	DB             string `mapstructure:"DB"`
	Port           string `mapstructure:"PORT"`
	DBName         string `mapstructure:"DB_NAME"`
	SessionsSecret string `mapstructure:"SESSIONS_SECRET"`
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
