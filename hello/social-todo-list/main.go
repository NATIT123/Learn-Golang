package main

import (
	"log"
	"main/cmd"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ////MongoDb
	// DB_MONGO := os.Getenv("DB_MONGO")
	// DB_MONGO = strings.Replace(DB_MONGO, "db_username", os.Getenv("DB_MONGO_USER"), 1)
	// DB_MONGO = strings.Replace(DB_MONGO, "<db_password>", os.Getenv("DB_MONGO_PASSWORD"), 1)
	// store := storagemongo.CreateMongo(DB_MONGO)
	// client := store.Client

	cmd.Excucte()

}
