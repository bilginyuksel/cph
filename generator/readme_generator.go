package generator

import (
	"fmt"
)

var tokens = map[string]bool{
	"export": true,
	"const": true,
	"let": true,
	"var": true,
	"default": true,
	"class": true,
	"function": true,
	"async": true,
	"extends": true,
	"implements": true,
	"abstract": true,
	"import":true,
	"constructor": true,
	"return": true,
	"//": true,
	"/*": true,
}

var symbols = map[string]bool {
	"(": true,
	")": true,
	"{": true,
	"}": true,
	"\"": true,
	"'": true,
	"`": true,
	":": true,
	";": true
}

func Parse(content string) *TSFile{
	return nil
}

type AccessSpecifier string
const (
	Private AccessSpecifier = "private"
	Protected AccessSpecifier = "protected"
	Public AccessSpecifier = "public"
	Default AccessSpecifier = "default"
)

type TSFile struct {
	Name string
	Imports map[string][]string
	Exports map[string][]string
	Functions map[string]Function
	Classes map[string]Class
	Variables map[string]Variable
	Interface map[string]Interface
	Enum map[string]Enum
	Comments []string
}

type Class struct {
	Abstract bool
	Export bool
	Default bool
	// Inner class...
	Name string
	Inherited *Class
	Implemented []Interface
	Annotations []Annotation
	Attributes []Attribute
	Methods []Method
	Constructors []Constructor
}

type Interface struct {
	Export bool
	Name string
	Inherited []Interface
	Variables map[string]string
	Methods []Method
}

type Enum struct {
	Name string
	Values map[string]string
}

type Constructor struct {
	AccessSpecifier AccessSpecifier
	Parameters []Parameter
	Parent *Class
}

type Method struct {
	Static bool
	Abstract bool
	Async bool
	Name string
	Annotations []Annotation
	Parameters []Parameter
	Return string
	AccessSpecifier AccessSpecifier
	Parent *Class
	DocString *DocString
}

type Function struct {
	Export bool
	Async bool
	Name string
	Annotations []Annotation
	Parameters []Parameter
	Return string
	Parent *File
	DocString *DocString
}

type Variable struct {
	Export bool
	Name string
	Type string
}

type Attribute struct {
	AccessSpecifier AccessSpecifier
	Name string
	Type string
}

type Annotation struct {
	Name string
	Param string
}

type Parameter struct {
	Name string
	Type string
	DefaultValue string
}

type ReturnDoc struct{
	Return string
	Description string
}

type DocString struct {
	Description string
	Params map[string]string
	ReturnDoc ReturnDoc
}
