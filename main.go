package main

import (
	"fmt"
	"go-action-runner/pkg/readme"
	"os"
	"strconv"

	"go-action-runner/github"
	"go-action-runner/gitlab"
	"go-action-runner/pkg/stats"
)

// Repo TODO
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
		fmt.Errorf("error occured during fetching GitHub Repos, %s", err)
	}

	gitlabRepos, err := gitlab.GetRepos(intProjectID)
	if err != nil {
		fmt.Errorf("error occured during fetching GitLab Repos, %s", err)
	}
	var exists bool
	exists = false
	t := stats.NewTable([]string{"Repo Name", "GitHub", "GitHub Org", "GitLab", "GitLab Project"})
	for _, githubRepo := range githubRepos {
		for _, gitlabRepo := range gitlabRepos {
			if githubRepo.Name == gitlabRepo.Name {
				exists = true
				t.AddRow([]string{githubRepo.Name, "<center>:white_check_mark:</center>", org, ":white_check_mark:", projectID})
			}
		}
		if !exists {
			t.AddRow([]string{githubRepo.Name, "<center>:x:</center>", org, ":white_check_mark:", projectID})
		}
	}
	//readme.Update("stats.md", t.String())
	updateStatus := readme.UpdateGitHubRepoFile([]byte(t.String()), os.Getenv("GH_STATS_REPO_NAME"), os.Getenv("GH_STATS_ORG_NAME"), "stats.md")
	exists = false
	fmt.Printf("The migration status has been updated on Stats file successfully with status %s.", updateStatus)
}
