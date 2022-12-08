package gitlab

import (
	"testing"
)

func TestGetRepos(t *testing.T) {
	projectID := 123456 // Replace with the ID of a GitLab project that you have access to
	repos, err := GetRepos(projectID)
	if err != nil {
		t.Errorf("GetRepos() returned an error: %v", err)
	}

	if len(repos) == 0 {
		t.Errorf("GetRepos() returned no repositories")
	}

	// Check that the returned repositories have the expected fields
	for _, repo := range repos {
		if repo.ID == 0 {
			t.Errorf("Repository ID is not set")
		}

		if repo.Name == "" {
			t.Errorf("Repository name is not set")
		}

		if repo.URL == "" {
			t.Errorf("Repository URL is not set")
		}
	}
}
