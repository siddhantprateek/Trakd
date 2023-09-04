package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type APIConfiguration struct {
	Port int
}

func GetEnv(ENV_VAR string) string {

	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load .env file")
	}
	envVar := os.Getenv(ENV_VAR)
	return envVar
}
