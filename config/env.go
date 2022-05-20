package config

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envMap, err := godotenv.Read()

	if err != nil {
		log.Fatal(err)
	}

	return envMap[key]

}
