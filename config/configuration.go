package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

type Configuration struct {
	Env    string
	Server struct {
		Host string
		Port int
	}
	Database struct {
		Dsn string
	}
	JWT struct {
		Secret         string
		ExpirationTime time.Duration
	}
}

var globalConfiguration *Configuration

func LoadConfiguration() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".") // path to look for the config file in
	viper.AutomaticEnv()     // read environment variables that match keys

	// Optional: replace dots in environment variables with underscores

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read configuration file")
	}

	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		return errors.Wrap(err, "failed to unmarshal configuration")
	}

	return nil
}

func GetGlobalConfig() *Configuration {
	return globalConfiguration
}
