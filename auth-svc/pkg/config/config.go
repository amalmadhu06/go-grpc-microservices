package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`           // Port represents the application port.
	DBUrl        string `mapstructure:"DB_URL"`         // DBUrl represents URL for connecting with database
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"` // JWTSecretKey represents the secret used to sign JSON Web Token used for authentication
}

var envs = []string{
	"PORT",
	"DB_URL",
	"JWT_SECRET_KEY",
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
