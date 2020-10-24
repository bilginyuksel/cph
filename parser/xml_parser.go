package parser

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// Plugin ...
type Plugin struct {
	XMLName      xml.Name   `xml:"plugin"`
	ID           string     `xml:"id,attr"`
	Version      string     `xml:"version,attr"`
	Xmlns        string     `xml:"xmlns,attr"`
	XmlnsAndroid string     `xml:"xmlns:android,attr"`
	Name         string     `xml:"name"`
	Description  string     `xml:"description"`
	License      string     `xml:"license"`
	Keywords     []string   `xml:"keywords"`
	Author       string     `xml:"author"`
	Engines      []Engine   `xml:"engines"`
	JsModule     []JSModule `xml:"js-module"`
	Platform     Platform   `xml:"platform"`
}

// Platform ...
type Platform struct {
	XMLName     xml.Name     `xml:"platform"`
	Name        string       `xml:"name,attr"`
	Hooks       []Hook       `xml:"hook"`
	ConfigFiles []ConfigFile `xml:"config-file"`
}

// ConfigFile ...
type ConfigFile struct {
	XMLName         xml.Name         `xml:"config-file"`
	Target          string           `xml:"target,attr"`
	Parent          string           `xml:"parent,attr"`
	UsesPermissions []UsesPermission `xml:"uses-permission"`
	Receivers       []Receiver       `xml:"receiver"`
	Metadata        Metadata         `xml:"meta-data"`
}

// Metadata ...
type Metadata struct {
	XMLName      xml.Name `xml:"meta-data"`
	AndroidName  string   `xml:"android:value,attr"`
	AndroidValue string   `xml:"android:name,attr"`
}

// Receiver ...
type Receiver struct {
	XMLName       xml.Name       `xml:"receiver"`
	AndroidName   string         `xml:"android:name,attr"`
	IntentFilters []IntentFilter `xml:"intent-filter"`
}

// IntentFilter ...
type IntentFilter struct {
	XMLName xml.Name `xml:"intent-filter"`
	Action  Action   `xml:"action"`
}

// Action ...
type Action struct {
	XMLName         xml.Name `xml:"action"`
	AndroidName     string   `xml:"android:name,attr"`
	AndroidEnabled  string   `xml:"android:enabled,attr"`
	AndroidExported string   `xml:"android:exported,attr"`
}

// UsesPermission ...
type UsesPermission struct {
	XMLName     xml.Name `xml:"uses-permission"`
	AndroidName string   `xml:"android:name,attr"`
}

// Hook ...
type Hook struct {
	XMLName xml.Name `xml:"hook"`
	Type    string   `xml:"type,attr"`
	Src     string   `xml:"src,attr"`
}

// Engine ...
type Engine struct {
	XMLName xml.Name `xml:"engine"`
	Name    string   `xml:"name,attr"`
	Version string   `xml:"version,attr"`
}

// SourceFile ...
type SourceFile struct {
	XMLName   xml.Name `xml:"source-file"`
	Src       string   `xml:"src,attr"`
	TargetDir string   `xml:"target-dir,attr"`
}

// JSModule ...
type JSModule struct {
	XMLName  xml.Name `xml:"js-module"`
	Name     string   `xml:"name,attr"`
	Src      string   `xml:"src,attr"`
	Clobbers Clobbers `xml:"clobbers"`
}

// Clobbers ...
type Clobbers struct {
	XMLName xml.Name `xml:"clobbers"`
	Target  string   `xml:"target,attr"`
}

// ParseXML ...
func ParseXML(filename string) (*Plugin, error) {

	if !checkIsAnXMLFile(filename) {
		return nil, errors.New("only XML files can be parsed with this method")
	}
	xmlFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var plugin Plugin
	xml.Unmarshal(byteValue, &plugin)
	return &plugin, nil
}

func checkIsAnXMLFile(filename string) bool {
	splittedString := strings.Split(filename, ".")
	return splittedString[len(splittedString)-1] == "xml"
}
