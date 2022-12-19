package main

import (
	"fmt"
	"migrator/pkg/readme"
	"os"
	"strconv"

	"migrator/cmd"
	"migrator/github"
	"migrator/gitlab"
	"migrator/pkg/stats"
)

func main() {
	cmd.Execute()

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

	t := stats.NewTable([]string{"Repo Name", "GitHub Org", "GitLab Project", "GitLab Status", "GitHub Status", "Migrated"})

	for _, gitlabRepo := range gitlabRepos {
		searched := false
		for _, githubRepo := range githubRepos {
			if githubRepo.Name == gitlabRepo.Name {
				r := stats.NewRow(gitlabRepo.Name, org, projectID, true)
				t.AddRow(*r)
				searched = true
				break
			}
		}
		if !searched {
			r := stats.NewRow(gitlabRepo.Name, org, projectID, false)
			t.AddRow(*r)
		}
	}
	readme.Update("stats.md", t.String())
	//updateStatus := readme.UpdateGitHubRepoFile([]byte(t.String()), "go-action-runner", "mouismail", "stats.md")
	//fmt.Println(updateStatus)
	// fmt.Printf("The migration status has been updated on Stats file successfully with status %s.", updateStatus)
	fmt.Printf("The migration status has been updated on Stats file successfully")
}
