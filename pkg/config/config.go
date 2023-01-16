package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
}

type Config interface {
	Get(key string) any
	GetString(key string) string
}

func LoadEnv() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Can't open .env file")
	}
	return &config{}
}

func (c *config) Get(key string) any {
	data := os.Getenv(key)
	return data
}

func (c *config) GetString(key string) string {
	data := c.Get(key).(string)
	return data
}
