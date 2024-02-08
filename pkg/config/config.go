package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (c *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return

}
