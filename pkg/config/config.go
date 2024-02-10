package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`

	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
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
