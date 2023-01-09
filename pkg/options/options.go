package options

import (
	"errors"
	"fmt"
	"migrator/build"
)

type Options struct {
	Source          build.SourceType
	SourceProjectId int    `json:"sourceProjectId"`
	GitHubOrg       string `json:"gitHubOrg"`
}

func NewOptions(source build.SourceType, githubOrg string, projectId int) *Options {
	return &Options{
		source,
		projectId,
		githubOrg,
	}
}

func (o *Options) AddOptions(source build.SourceType, githubOrg string, projectId int) {
	o.Source = source
	o.SourceProjectId = projectId
	o.GitHubOrg = githubOrg
}

func (o *Options) Check() error {
	fmt.Println(o.Source.Source)
	if o.Source.Source == "GitLab" && o.Source.HasProject {
		if o.SourceProjectId == 0 {
			return errors.New("project ID can't be empty for GitLab")
		}
	}
	return nil
}

func Generate(o *Options) {
	fmt.Println(o.Source)
	fmt.Println(o.GitHubOrg)
	fmt.Println(o.SourceProjectId)
}
