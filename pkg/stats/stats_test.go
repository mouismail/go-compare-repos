package stats

import (
	"reflect"
	"testing"
)

func TestNewTable(t *testing.T) {
	type args struct {
		headers []string
	}
	var tests []struct {
		name string
		args args
		want *Table
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTable(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_AddRow(t1 *testing.T) {
	type fields struct {
		headers []string
		rows    [][]string
	}
	type args struct {
		row []string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				headers: tt.fields.headers,
				rows:    tt.fields.rows,
			}
			t.AddRow(tt.args.row)
		})
	}
}

func TestTable_String(t1 *testing.T) {
	type fields struct {
		headers []string
		rows    [][]string
	}
	var tests []struct {
		name   string
		fields fields
		want   string
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				headers: tt.fields.headers,
				rows:    tt.fields.rows,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
