package parser

import "testing"

func TestParseVariables_tokenizeVar1(t *testing.T) {
	content := `const foo  ="Life \"is all about quotes\"";let l;     var baz:string; const koz:   number   =     5;var soz; var loz = 'hi'`

	expected := []string{"const", "foo", "=", "Life \"is all about quotes\"", ";", "let", "l", ";", "var", "baz", ":", "string", ";", "const", "koz", ":", "number", "=", "5", ";", "var", "soz", ";", "var", "loz", "=", "hi"}
	given := ParseVariables(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i:=0; i<len(expected); i++{
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}

func TestParseVariables_tokenizeVar2(t *testing.T) {
	content := `const f='\'\'hello\'\'';let l="\"\"world\"\"";`

	expected := []string{"const","f","=","''hello''", ";", "let", "l", "=", "\"\"world\"\"", ";"}

	given := ParseVariables(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i:=0; i<len(expected); i++{
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}
/*
func TestParseVariables_tokenizeVar3(t *testing.T) {
	content := `const f: string="hello world"
				const l="wow"
let yey; let how: string
var v
var l = "lo\"l\"ol";let l;`

	expected := []string{"const", "f", ":", "string", "=", "hello world", "const", "l", "=", "wow", "let",
		"yey", ";", "let", "how", ":", "string", "var", "v", "var", "l", "=","lo\"l\"ol", ";", "let", "l", ";"}

	given := ParseVariables(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i:=0; i<len(expected); i++{
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}

}
*/
