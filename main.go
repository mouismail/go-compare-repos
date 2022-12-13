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

type MigrationStatus struct {
	Repo       string `json:"name"`
	IsMigrated bool   `json:"bool"`
}

func NewMigrationStatus(repo string, isMigrated bool) *MigrationStatus {
	return &MigrationStatus{
		repo,
		isMigrated,
	}
}

func (m *MigrationStatus) AddMigrationStatus(repo string, isMigrated bool) {
	m.Repo = repo
	m.IsMigrated = isMigrated
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

	t := stats.NewTable([]string{"Repo Name", "GitHub Org", "GitLab Project", "GitHub Status", "GitLab Status", "Migrated"})

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
	// updateStatus := readme.UpdateGitHubRepoFile([]byte(t.String()), "go-action-runner", "mouismail", "stats.md")
	// fmt.Printf("The migration status has been updated on Stats file successfully with status %s.", updateStatus)
	fmt.Printf("The migration status has been updated on Stats file successfully")
}

//
//func compareGitHubToGitLabRepos(githubRepo string, gitlabRepos []gitlab.Repo, t *stats.Table, org, projectID string) {
//	for _, gitLabRepo := range gitlabRepos {
//		if gitLabRepo.Name == githubRepo {
//			t.AddRow([]string{githubRepo, ":white_check_mark:", org, ":white_check_mark:", projectID})
//		}
//	}
//	t.AddRow([]string{githubRepo, ":x:", org, ":white_check_mark:", projectID})
//}
//
//func compareGitLabToGitHubRepos(gitLabRepo string, githubRepos []github.Repo, t *stats.Table, org, projectID string) {
//	for _, githubRepo := range githubRepos {
//		if githubRepo.Name != gitLabRepo {
//			t.AddRow([]string{gitLabRepo, ":x:", org, ":white_check_mark:", projectID})
//		}
//	}
//}
