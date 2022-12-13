package gitlab

import (
	"reflect"
	"testing"
)

func TestGetRepos(t *testing.T) {
	type args struct {
		projectID int
	}
	var tests []struct {
		name    string
		args    args
		want    []Repo
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRepos(tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepos() got = %v, want %v", got, tt.want)
			}
		})
	}
}
