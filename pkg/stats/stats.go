package stats

import (
	"bytes"
	"fmt"
)

type Table struct {
	headers []string
	rows    [][]string
}

func NewTable(headers []string) *Table {
	return &Table{
		headers: headers,
		rows:    [][]string{},
	}
}

// AddRow adds a row to the table.
func (t *Table) AddRow(row []string) {
	t.rows = append(t.rows, row)
}

func (t *Table) String() string {
	var b bytes.Buffer

	// write headers
	for _, h := range t.headers {
		b.WriteString(fmt.Sprintf("| %s ", h))
	}
	b.WriteString("|\n")

	// write separator
	for range t.headers {
		b.WriteString("| :---: ")
	}
	b.WriteString("|\n")

	// write rows
	for _, r := range t.rows {
		for _, c := range r {
			b.WriteString(fmt.Sprintf("| %s ", c))
		}
		b.WriteString("|\n")
	}

	return b.String()
}
