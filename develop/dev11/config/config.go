package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP
}

type HTTP struct {
	host         string
	port         string
	timeout      string
	readTimeout  string
	writeTimeout string
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
		c.port = httpPort
	}

	httpHost := os.Getenv("HTTP_HOST")
	if httpHost != "" {
		c.host = httpHost
	}

	httpTimeout := os.Getenv("HTTP_TIMEOUT")
	if httpTimeout != "" {
		c.timeout = httpTimeout
	}

	httpReadTimeout := os.Getenv("HTTP_READTIMEOUT")
	if httpReadTimeout != "" {
		c.readTimeout = httpReadTimeout
	}

	httpWriteTimeout := os.Getenv("HTTP_WRITEIMEOUT")
	if httpWriteTimeout != "" {
		c.writeTimeout = httpWriteTimeout
	}

	return &c, nil
}
