package readme

import "testing"

func TestRead(t *testing.T) {
	type args struct {
		filePath string
	}
	var tests []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		filePath string
		content  string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Update(tt.args.filePath, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateGitHubRepoFile(t *testing.T) {
	type args struct {
		fileContents []byte
		repo         string
		org          string
		filePath     string
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateGitHubRepoFile(tt.args.fileContents, tt.args.repo, tt.args.org, tt.args.filePath); got != tt.want {
				t.Errorf("UpdateGitHubRepoFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
