package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Environmental .env not found!")
	}
	apiKey := os.Getenv("FREELANCER_API_KEY")
	if apiKey == "" {
		log.Fatal("FREELANCER_API_KEY environment variable not set")
	}
	projects, err := FetchProjects(apiKey)
	if err != nil {
		log.Fatal("Error fetching projects", err)
	}
	fmt.Println(len(projects))
}
