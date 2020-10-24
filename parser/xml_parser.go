package parser

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

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

// CreateXML ...
func CreateXML(plugin *Plugin, filename string) error {

	file, _ := xml.MarshalIndent(plugin, "", "\t")
	file = []byte(xml.Header + string(file))
	error := ioutil.WriteFile(filename, file, 0644)
	return error
}

// Plugin ...
type Plugin struct {
	ID           string     `xml:"id,attr"`
	Version      string     `xml:"version,attr"`
	Xmlns        string     `xml:"xmlns,attr"`
	XmlnsAndroid string     `xml:"android,attr"`
	Name         string     `xml:"name"`
	Description  string     `xml:"description"`
	License      string     `xml:"license,"`
	Keywords     string     `xml:"keywords"`
	Author       string     `xml:"author,omitempty"`
	Engines      *Engines   `xml:"engines,omitempty"`
	JsModule     []JSModule `xml:"js-module"`
	Platform     Platform   `xml:"platform"`
}

// Engines ...
type Engines struct {
	Engine []Engine `xml:"engine"`
}

// Platform ...
type Platform struct {
	Name        string       `xml:"name,attr"`
	Hooks       []Hook       `xml:"hook"`
	ConfigFiles []ConfigFile `xml:"config-file"`
	Frameworks  []Framework  `xml:"framework"`
	SourceFiles []SourceFile `xml:"source-file"`
}

// Framework ...
type Framework struct {
	Src    string `xml:"src,attr"`
	Custom bool   `xml:"custom,attr,omitempty"`
	Type   string `xml:"type,attr,omitempty"`
}

// ConfigFile ...
type ConfigFile struct {
	Target          string           `xml:"target,attr,omitempty"`
	Parent          string           `xml:"parent,attr,omitempty"`
	UsesPermissions []UsesPermission `xml:"uses-permission"`
	Receivers       []Receiver       `xml:"receiver"`
	Metadata        *Metadata        `xml:"meta-data,omitempty"`
	Services        []Service        `xml:"service"`
	Features        []Feature        `xml:"feature"`
}

// Feature ...
type Feature struct {
	Name   string  `xml:"name,attr,omitempty"`
	Params []Param `xml:"param"`
}

// Param ...
type Param struct {
	Name  string `xml:"name,attr,omitempty"`
	Value string `xml:"value,attr,omitempty"`
}

// Service ...
type Service struct {
	AndroidName     string         `xml:"name,attr,omitempty"`
	AndroidExported bool           `xml:"exported,attr,omitempty"`
	IntentFilters   []IntentFilter `xml:"intent-filter"`
}

// Metadata ...
type Metadata struct {
	AndroidName  string `xml:"value,attr,omitempty"`
	AndroidValue string `xml:"name,attr,omitempty"`
}

// Receiver ...
type Receiver struct {
	AndroidName     string         `xml:"name,attr,omitempty"`
	AndroidEnabled  bool           `xml:"enabled,attr,omitempty"`
	AndroidExported bool           `xml:"exported,attr,omitempty"`
	IntentFilters   []IntentFilter `xml:"intent-filter"`
}

// IntentFilter ...
type IntentFilter struct {
	Actions    []Action   `xml:"action"`
	Categories []Category `xml:"category"`
	Datas      []Data     `xml:"data"`
}

// Category ...
type Category struct {
	AndroidName string `xml:"name,attr"`
}

// Data ...
type Data struct {
	AndroidScheme string `xml:"scheme,attr"`
}

// Action ...
type Action struct {
	AndroidName string `xml:"name,attr"`
}

// UsesPermission ...
type UsesPermission struct {
	AndroidName string `xml:"name,attr"`
}

// Hook ...
type Hook struct {
	Type string `xml:"type,attr"`
	Src  string `xml:"src,attr"`
}

// Engine ...
type Engine struct {
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr,utf-8"`
}

// SourceFile ...
type SourceFile struct {
	Src       string `xml:"src,attr"`
	TargetDir string `xml:"target-dir,attr"`
}

// JSModule ...
type JSModule struct {
	Name     string    `xml:"name,attr"`
	Src      string    `xml:"src,attr"`
	Clobbers *Clobbers `xml:"clobbers,omitempty"`
}

// Clobbers ...
type Clobbers struct {
	Target string `xml:"target,attr,omitempty"`
}

func checkIsAnXMLFile(filename string) bool {
	splittedString := strings.Split(filename, ".")
	return splittedString[len(splittedString)-1] == "xml"
}
