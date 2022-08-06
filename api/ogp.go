package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Print() {
	fmt.Println(os.Getenv("NAME"), os.Getenv("EDITOR"))
}
