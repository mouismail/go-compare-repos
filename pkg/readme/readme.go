package readme

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func Update(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func Read(filePath string) (string, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file's content
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func UpdateGitHubRepoFile(fileContents []byte, repo string, org string, filePath string) string {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GHEC_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	commit := &github.RepositoryContentFileOptions{
		Message: github.String("Update stats.md file"),
		Content: fileContents,
		Branch:  github.String("main"),
	}
	_, resp, err := client.Repositories.UpdateFile(ctx, org, repo, filePath, commit)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return resp.Status
}
