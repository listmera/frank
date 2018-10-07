package env

import (
	"github.com/joho/godotenv"
	"github.com/listmera/frank/utils"
	"os"
)

func SetEnv () {
	currEnv := GetOr("GO_ENV", "development")
	if currEnv == "development" {
		err := godotenv.Load()
		utils.CheckErr(err)
	}
}

func GetOr (key string, or string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return or
	}
	return value
}