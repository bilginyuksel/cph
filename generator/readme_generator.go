package generator

import (
	"fmt"
)

func main(){
	fmt.Println("Generator package main method.")
}


type TSFile struct {
	Functions map[string]Function
}

type AccessSpecifier string
const (
	Private AccessSpecifier = "private"
	Protected AccessSpecifier = "protected"
	Public AccessSpecifier = "public"
	Default AccessSpecifier = "default"
)

type Method struct {
	Static bool
	Annotations []Annotation
	Parameters []Parameter
	ReturnType DataType
	AccessSpecifier AccessSpecifier
	Parent *Class
	DocString *DocString
}

type Function struct {
	Export bool
	Annotations []Annotation
	Parameters []Parameter
	ReturnType DataType
	Parent *File
	DocString *DocString
}

type DocString struct {}
type DataType struct {}
type Function struct {}
type Class struct {}
type Annotation struct {}
type Parameter struct {}
