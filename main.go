package main

import (
	"fmt"
	"go-action-runner/github"
	"go-action-runner/gitlab"
	"os"
	"strconv"
)

type Repo struct {
	Name         string `json:"name"`
	Platform     string `json:"platform"`
	Organization string `json:"organization"`
	Project      string `json:"project"`
	Type         string `json:"type"`
}

func main() {
	org := os.Getenv("GH_ORG_NAME")
	projectID := os.Getenv("GL_PROJECT_ID")
	intProjectID, _ := strconv.Atoi(projectID)

	githubRepos, err := github.GetRepos(org)
	if err != nil {
		// Handle the error
	}

	gitlabRepos, err := gitlab.GetRepos(intProjectID)
	if err != nil {
		// Handle the error
	}

	for _, githubRepo := range githubRepos {
		for _, gitlabRepo := range gitlabRepos {
			if githubRepo.Name == gitlabRepo.Name {
				fmt.Printf("%s found in GitHub", gitlabRepo.Name)
			}
		}
	}
}
