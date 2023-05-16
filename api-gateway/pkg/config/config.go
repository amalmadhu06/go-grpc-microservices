package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config represents the application configuration.
// It holds various settings that can be loaded from environment variables or a configuration file.
type Config struct {
	Port          string `mapstructure:"PORT"`            // Port represents the application port.
	AuthSuvUrl    string `mapstructure:"AUTH_SUV_URL"`    // AuthSuvUrl represents the URL for the authentication service.
	ProductSuvUrl string `mapstructure:"PRODUCT_SUV_URL"` // ProductSuvUrl represents the URL for the product service.
	OrderSuvUrl   string `mapstructure:"ORDER_SUV_URL"`   // OrderSuvUrl represents the URL for the order service.
}

var envs = []string{
	"PORT",
	"AUTH_SUV_URL",
	"PRODUCT_SUV_URL",
	"ORDER_SUV_URL",
}

// LoadConfig loads the application configuration from environment variables or a configuration file.
func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
