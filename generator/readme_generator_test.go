package generator

import (
	"testing"

	"github.com/bilginyuksel/cph/parser"
)

func TestCreateTableEnum_ReturnEnumTable(t *testing.T) {
	var enum = parser.Enum{Name: "Colors",
		Items: []parser.EnumItem{
			parser.EnumItem{Name: "RED", Value: "-1"},
			parser.EnumItem{Name: "GREEN", Value: "2"},
			parser.EnumItem{Name: "BLUE", Value: "5"},
		}}
	given := createTableEnum(&enum)
	expected := &TableEnum{Name: "Colors", Description: "-",
		Rows: []EnumRow{
			EnumRow{"RED", "-1", "-"},
			EnumRow{"GREEN", "2", "-"},
			EnumRow{"BLUE", "5", "-"},
		}}
	if !compareEnumTables(given, expected) {
		t.Errorf("given= %v, expected= %v", given, expected)
	}
}

func TestConvertEnumRow_ExpectStringRow(t *testing.T) {
	row := EnumRow{Name: "RED", Value: "-1", Description: "-"}
	expected := "|RED|-1|-|"
	given := row.toSTR()
	if given != expected {
		t.Error()
	}
}

func TestTableToSTR_ReturnSTRTable(t *testing.T) {
	tenum := TableEnum{Name: "Colors", Description: "-",
		Rows: []EnumRow{
			EnumRow{"RED", "-1", "-"},
			EnumRow{"GREEN", "2", "-"},
			EnumRow{"BLUE", "5", "-"},
		}}
	given := tenum.toSTR(3)
	expected := `### Colors
-
|Name|Value|Description|
|---|---|---|
|RED|-1|-|
|GREEN|2|-|
|BLUE|5|-|
`
	if given != expected {
		t.Errorf("given= %v, expected= %v", given, expected)
	}
}

func TestTableToSTR_ReturnSTRTable2(t *testing.T) {
	tenum := TableEnum{Name: "Colors", Description: "-",
		Rows: []EnumRow{
			EnumRow{"RED", "-1", "-"},
			EnumRow{"GREEN", "2", "-"},
			EnumRow{"BLUE", "5", "-"},
		}}
	given := tenum.toSTR(3)
	expected := `### Colors
-
|Name|Value|Description|
|---|---|---|
|RED|-1|-|
|GREEN|2|-|
|BLUE|5|-|
`

	tenum2 := TableEnum{Name: "Event", Description: "-",
		Rows: []EnumRow{
			EnumRow{"MY_EVENT", "my_best_event_1", "-"},
			EnumRow{"MY_EVENT2", "my_best_event_2", "-"},
		}}
	given2 := tenum2.toSTR(3)
	expected2 := `### Event
-
|Name|Value|Description|
|---|---|---|
|MY_EVENT|my_best_event_1|-|
|MY_EVENT2|my_best_event_2|-|
`
	if given2 != expected2 {
		t.Errorf("given= %v, expected= %v", given, expected)
	}
}

func compareEnumTables(given *TableEnum, expected *TableEnum) bool {
	if given.Name != expected.Name {
		return false
	}
	if given.Description != expected.Description {
		return false
	}
	if !compareRows(given.Rows, expected.Rows) {
		return false
	}
	return true
}

func compareRows(given []EnumRow, expected []EnumRow) bool {
	if len(expected) != len(given) {
		return false
	}
	for i := 0; i < len(expected); i++ {
		if given[i].Name != expected[i].Name || given[i].Value != expected[i].Value || given[i].Description != expected[i].Description {
			return false
		}
	}
	return true
}
