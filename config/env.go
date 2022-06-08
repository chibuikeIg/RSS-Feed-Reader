package config

import (
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {

	err := godotenv.Load()

	if err != nil {
		Log(err.Error())
	}

	envMap, err := godotenv.Read()

	if err != nil {
		Log(err.Error())
	}

	return envMap[key]

}
