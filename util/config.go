package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of application
// The values are read by viper from a config file or env file
type Config struct {
	DBDriver 						string				`mapstructure:"DB_DRIVER"`
	DBSource 						string				`mapstructure:"DB_Source"`
	ServerAddress 			string				`mapstructure:"SERVER_ADDRESS"`
	SymmetricTokenKey		string				`mapstructure:"SYMMETRIC_TOKEN_KEY"`
	AccessTokenDuration	time.Duration	`mapstructure:"ACCESS_TOKEN_DURATION"`
}

// LoadConfig reads configuration from config file or env file
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