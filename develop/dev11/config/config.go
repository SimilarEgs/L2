package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP
}

type HTTP struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ParseConfig func - returns parsed config struct with filled fileds, which was read from the config file
func ParseConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.yaml")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("[Error] .env file didn't load: %s", err.Error())
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config

	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort != "" {
		c.Port = httpPort
	}

	httpHost := os.Getenv("HTTP_HOST")
	if httpHost != "" {
		c.Host = httpHost
	}

	httpReadTimeout := os.Getenv("HTTP_READTIMEOUT")
	if httpReadTimeout != "" {
		c.ReadTimeout, err = time.ParseDuration(httpReadTimeout)
		if err != nil {
			return nil, err
		}
	}

	httpWriteTimeout := os.Getenv("HTTP_WRITEIMEOUT")
	if httpWriteTimeout != "" {
		c.WriteTimeout, err = time.ParseDuration(httpWriteTimeout)
		if err != nil {
			return nil, err
		}
	}
	return &c, nil
}
