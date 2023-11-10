package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
		fmt.Printf("loaded env vars - PORT=%s",os.Getenv("PORT"))
	}
}
