package stats

import (
	"bytes"
	"fmt"
)

type Table struct {
	headers []string
	rows    []TableRow
}

type TableRow struct {
	GitHubRepo      string `json:"gitHubRepo"`
	GitHubOrg       string `json:"gitHubOrg"`
	GitLabRepo      string `json:"gitLabRepo"`
	GitLabProjectId string `json:"gitLabProjectId"`
	IsExists        bool   `json:"isExists"`
}

func NewRow(repo, org, projectId string, isExist bool) *TableRow {
	return &TableRow{
		GitHubRepo:      repo,
		GitHubOrg:       org,
		GitLabRepo:      repo,
		GitLabProjectId: projectId,
		IsExists:        isExist,
	}
}

func NewTable(headers []string) *Table {
	return &Table{
		headers: headers,
		rows:    []TableRow{},
	}
}

func (t *Table) AddRow(row TableRow) {
	t.rows = append(t.rows, row)
}

func (t *Table) String() string {
	var b bytes.Buffer

	for _, h := range t.headers {
		b.WriteString(fmt.Sprintf("| %s ", h))
	}
	b.WriteString("|\n")

	for range t.headers {
		b.WriteString("| :---: ")
	}
	b.WriteString("|\n")

	for _, r := range t.rows {

		b.WriteString(fmt.Sprintf("| %s ", r.GitHubRepo))
		b.WriteString(fmt.Sprintf("| %s ", r.GitHubOrg))
		b.WriteString(fmt.Sprintf("| %s ", r.GitLabProjectId))
		if r.IsExists {
			b.WriteString(fmt.Sprintf("| :white_check_mark: "))
			b.WriteString(fmt.Sprintf("| :white_check_mark: "))
			b.WriteString(fmt.Sprintf("| :ok_hand: "))
		} else {
			b.WriteString(fmt.Sprintf("| :white_check_mark: "))
			b.WriteString(fmt.Sprintf("| :x: "))
			b.WriteString(fmt.Sprintf("| :thumbsdown: "))
		}
		b.WriteString("|\n")
	}

	return b.String()
}
