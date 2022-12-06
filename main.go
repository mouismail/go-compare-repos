package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Repo is a struct that represents a repository.
type Repo struct {
	Name         string `json:"name"`
	Platform     string `json:"platform"`
	Organization string `json:"organization"`
	Project      string `json:"project"`
	Type         string `json:"type"`
}

// CompareRepos compares the repositories in the specified organization with the repositories in the provided JSON file.
func CompareRepos(org string, file string) error {
	// Create a new OAuth client using the GITHUB_TOKEN secret.
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GHEC_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	// Get a list of all the repositories in the organization.
	repos, _, err := client.Repositories.List(context.Background(), org, nil)
	if err != nil {
		return err
	}

	// Open the provided JSON file and read its contents.
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// Unmarshal the JSON file contents into a slice of Repo structs.
	var reposFromFile []Repo
	if err := json.Unmarshal(fileContents, &reposFromFile); err != nil {
		return err
	}

	// Compare the repositories in the organization with the repositories in the JSON file.
	for _, repo := range repos {
		found := false
		for _, repoFromFile := range reposFromFile {
			if repo.GetName() == repoFromFile.Name {
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Repository %s is not in the file.\n", repo.GetName())
		}
	}

	return nil
}

func main() {
	// Set the organization and file inputs.
	org := os.Getenv("org")
	file := os.Getenv("file")

	// Compare the repositories in the organization with the repositories in the file.
	err := CompareRepos(org, file)
	if err != nil {
		log.Fatal(err)
	}
}
