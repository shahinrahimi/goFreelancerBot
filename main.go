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
	focusedSkills := []string{
		"react",
		"go",
		"javascript",
	}
	focusedProjects, err := FetchProjects(apiKey, focusedSkills)
	if err != nil {
		log.Fatal("Error fetching projects", err)
	}
	// other skills
	otherSkills := []string{
		"vba",
		"python",
		"libreoffice",
		"excel",
		"shiny",
		"scrapping",
	}
	otherProjects, err := FetchProjects(apiKey, otherSkills)
	if err != nil {
		log.Fatal("Error fetching projects", err)
	}

	for index, project := range otherProjects {
		fmt.Println(index, project.Title)
	}
	ownerIDs := []int{}
	for index, project := range focusedProjects {
		ownerIDs = append(ownerIDs, project.OwnerID)
		fmt.Println(index, project.OwnerID, project.Budget.Maximum, project.Title)
	}

	owners, err := FetchOwners(apiKey, ownerIDs)
	if err != nil {
		log.Fatal("Error fetching owners", err)
	}

	for _, owner := range owners {
		fmt.Println(owner.ID, owner.Username, owner.Location.Country.Name)
	}
	fmt.Println(len(owners))
	fmt.Println(len(focusedProjects), len(otherProjects))
	fmt.Println(CountCommonIDs(focusedProjects, otherProjects))

}
