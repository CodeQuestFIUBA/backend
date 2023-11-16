package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const EnvFilePath = ".env"

func EnvMongoURI() string {
	err := godotenv.Load(EnvFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_DB_URI")
}

func EnvDBName() string {
	err := godotenv.Load(EnvFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_DB_NAME")
}
