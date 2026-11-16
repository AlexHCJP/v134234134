package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	return &Config{
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}
}

func (config *Config) Addr() string {
	fmt.Println(config.Host + ":" + config.Port)
	return config.Host + ":" + config.Port
}

var config = NewConfig()
