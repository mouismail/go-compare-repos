package options

import (
	"migrator/build"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		o *Options
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Generate(tt.args.o)
		})
	}
}

func TestNewOptions(t *testing.T) {
	type args struct {
		source    build.SourceType
		githubOrg string
		projectId int
	}
	var tests []struct {
		name string
		args args
		want *Options
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOptions(tt.args.source, tt.args.githubOrg, tt.args.projectId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_AddOptions(t *testing.T) {
	type fields struct {
		Source          build.SourceType
		SourceProjectId int
		GitHubOrg       string
	}
	type args struct {
		source    build.SourceType
		githubOrg string
		projectId int
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				Source:          tt.fields.Source,
				SourceProjectId: tt.fields.SourceProjectId,
				GitHubOrg:       tt.fields.GitHubOrg,
			}
			o.AddOptions(tt.args.source, tt.args.githubOrg, tt.args.projectId)
		})
	}
}

func TestOptions_Check(t *testing.T) {
	type fields struct {
		Source          build.SourceType
		SourceProjectId int
		GitHubOrg       string
	}
	var tests []struct {
		name    string
		fields  fields
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				Source:          tt.fields.Source,
				SourceProjectId: tt.fields.SourceProjectId,
				GitHubOrg:       tt.fields.GitHubOrg,
			}
			if err := o.Check(); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
