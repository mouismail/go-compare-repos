# Migration to GitHub Repos count

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg?style=for-the-badge&logo=go)](https://github.com/gomods/athens)
[![GitHub license](https://img.shields.io/github/license/mouismail/go-compare-repos.svg?style=for-the-badge)](https://github.com/mouismail/go-compare-repos/blob/main/LICENSE.md)

[![Latest release](https://badgen.net/github/release/mouismail/go-compare-repos?style=for-the-badge&logo=appveyor)](https://github.com/mouismail/go-compare-repos/releases)


## Description

This repository contains a GitHub Action that compares repositories in GitHub with repositories in GitLab or Bitbucket.

## Inputs

### Environment Variables

- `GH_ORG_NAME` - The name of the GitHub organization
- `GL_PROJECT_ID` - The ID of the GitLab project
- `GHEC_ACCESS_TOKEN` - The access token for the GitHub Enterprise Cloud instance
- `GITLAB_ACCESS_TOKEN` - The access token for the GitLab instance
- `GH_STATS_REPO_NAME` - The repo where the stats file will be created
- `GH_STATS_ORG_NAME` - The organization where the stats file will be created


## Development

### Requirements

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)
- [Goreleaser](https://goreleaser.com/)
- [GitHub CLI](https://cli.github.com/)

### Build

```bash 
make build
```

### Test

```bash
make test
```

### Release

```bash
make release
```

### Run locally

```bash
make run
```


## Usage

### Local


```bash
export GH_ORG_NAME=demo-org
export GL_PROJECT_ID=123456
export GHEC_ACCESS_TOKEN=gh_access_token
export GITLAB_ACCESS_TOKEN=gl_access_token
```

```bash
go run main.go
```

### Example workflow

```yaml
name: Go

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: mouismail/go-compare-repos@v1
        with:
          go-version: '1.19'
          go-command: 'go build'
```
