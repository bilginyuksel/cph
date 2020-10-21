package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Cordova plugin.xml
type Plugin struct {
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Author      string   `xml:"author"`
	Keywords    []string `xml:"keywords"`
	License     string   `xml:"license"`
	JsModule    JSModule `xml:"js-module"`
	SourceFile  SourceFile
}

type SourceFile struct {
	Src       string
	TargetDir string
}
type JSModule struct {
	clobbers string `xml:"clobbers"`
}

// -------------
type Users struct {
	XMLName xml.Name `xml::"users"`
	Users   []User   `xml:"user"`
}
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func Run() {
	xmlFile, err := os.Open("parser/test.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened test.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var users Users
	xml.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}
