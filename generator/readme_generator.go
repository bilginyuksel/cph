package generator

import (
	"fmt"
	"strings"

	"github.com/bilginyuksel/cph/parser"
)

// HeaderRow ...
type HeaderRow struct {
	field1 string
	field2 string
	field3 string
}

var ClassHeader = HeaderRow{"Name", "Type", "Description"}
var EnumHeader = HeaderRow{"Name", "Value", "Description"}
var InterfaceHeader = HeaderRow{"Name", "Type", "Description"}

// TableEnum ...
type TableEnum struct {
	Name        string
	Description string
	Rows        []EnumRow
}

// EnumRow ...
type EnumRow struct {
	Name        string
	Value       string
	Description string
}

func createTableEnum(content *parser.Enum) *TableEnum {
	rows := []EnumRow{}
	for _, item := range content.Items {
		rows = append(rows, EnumRow{Name: item.Name, Value: item.Value, Description: "-"})
	}
	return &TableEnum{Name: content.Name, Description: "-", Rows: rows}
}

func (row EnumRow) toSTR() string {
	return fmt.Sprintf("|%s|%s|%s|", row.Name, row.Value, row.Description)
}

func (table TableEnum) toSTR(level int) string {
	title := strings.Repeat("#", level) + " " + table.Name
	desc := table.Description
	rtable := fmt.Sprintf("|%s|%s|%s|\n|---|---|---|\n", EnumHeader.field1, EnumHeader.field2, EnumHeader.field3)
	for _, row := range table.Rows {
		rtable += row.toSTR() + "\n"
	}
	return title + "\n" + desc + "\n" + rtable
}
