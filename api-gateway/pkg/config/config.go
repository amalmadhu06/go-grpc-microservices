package config

import "github.com/spf13/viper"

// Config represents the application configuration.
// It holds various settings that can be loaded from environment variables or a configuration file.
type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSuvUrl    string `mapstructure:"AUTH_SUV_URL"`
	ProductSuvUrl string `mapstructure:"PRODUCT_SUV_URL"`
	OrderSuvUrl   string `mapstructure:"ORDER_SUV_URL"`
}

// LoadConfig loads the configuration settings.
// It reads the configuration file, sets up environment variable support, and unmarshal the settings into a Config struct.
// It returns the loaded Config and any error encountered during the process.
func LoadConfig() (c Config, err error) {

	// AddConfigPath adds the directory where the configuration file is located.
	viper.AddConfigPath("./pkg/config/envs")

	// SetConfigName sets the name of the configuration file to be read.
	viper.SetConfigName("dev")

	// SetConfigType sets the type of the configuration file.
	viper.SetConfigType("env")

	// AutomaticEnv enables automatic binding of environment variables to configuration values.
	viper.AutomaticEnv()

	// ReadInConfig reads the configuration file with the specified name and type.
	err = viper.ReadInConfig()

	// Check if there was an error reading the configuration file.
	if err != nil {
		return
	}

	// Unmarshal reads the configuration settings into the Config struct.
	err = viper.Unmarshal(&c)
	return
}
