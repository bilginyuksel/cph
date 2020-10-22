package parser

import (
	"fmt"
	"os"
)

// Cordova plugin.xml
type Plugin struct {
	ID           string   `xml:"id, Attr"`
	Version      string   `xml:"version, Attr"`
	Xmlns        string   `xml:"xmlns, Attr"`
	XmlnsAndroid string   `xml:"xmlns:android, Attr"`
	Name         string   `xml:"name"`
	Description  string   `xml:"description"`
	License      string   `xml:"license"`
	Keywords     []string `xml:"keywords"`
	Author       string   `xml:"author"`
	Engines      []Engine `xml:"engines"`
	JsModule     JSModule `xml:"js-module"`
	Platform     Platform `xml:"platform"`
}

type Platform struct {
}

type Engine struct {
}

type SourceFile struct {
	Src       string `xml:"src,Attr"`
	TargetDir string
}
type JSModule struct {
	clobbers string `xml:"clobbers"`
}

func Run() {
	xmlFile, err := os.Open("parser/test.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened test.xml")
	defer xmlFile.Close()

	// byteValue, _ := ioutil.ReadAll(xmlFile)

	// var users interface{}
	// xml.Unmarshal(byteValue, &users)

	// for i := 0; i < len(users.Users); i++ {
	// 	fmt.Println("User Type: " + users.Users[i].Type)
	// 	fmt.Println("User Name: " + users.Users[i].Name)
	// 	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	// }
}
