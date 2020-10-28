package main

import (
	"fmt"
	"path/filepath"

	"github.com/bilginyuksel/cordova-plugin-helper/parser"
	"github.com/bilginyuksel/cordova-plugin-helper/reader"
	"github.com/bilginyuksel/cordova-plugin-helper/writer"

	"github.com/alecthomas/kong"
)

func main() {
	prepareCliParser()
}

func prepareCliParser() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}

// SyncPluginXML ...
func SyncPluginXML(path string) error {
	if path == "" {
		path = "."
	}

	plugin, err := parser.ParseXML(fmt.Sprintf("%s/plugin.xml", path))
	if err != nil {
		return err
	}
	sourceFiles, _ := reader.FilePathWalkDir("src")
	plugin.Platform.NewSourceFrom(sourceFiles)
	jsModules, _ := reader.FilePathWalkDir("www")
	plugin.NewJsModulesFrom(jsModules)

	err = parser.CreateXML(plugin, "plugin.xml")
	return err
}

// AddLicenceTo ...
func AddLicenceTo(path string, extension string, licence string) error {
	if len(licence) == 0 {
		licence = "writer/licence"
	}
	files, err := reader.FilePathWalkDir(path)
	if err != nil {
		return err
	}

	for _, p := range files {
		ext := filepath.Ext(p)
		if ext == extension || len(extension) == 0 {
			writer.WriteLicenceToFile(p, licence)
		}
	}

	return nil
}

// Run ...
func (pl *PluginXMLCmd) Run(ctx *Context) error {
	return SyncPluginXML(pl.Path)
}

// Run ...
func (l *AddLicenseCmd) Run(ctx *Context) error {
	return AddLicenceTo(l.Path, l.Extension, l.License)
}

// Context ...
type Context struct {
	Debug bool
}

// PluginXMLCmd ...
type PluginXMLCmd struct {
	Path string `help:"The folder path to sync."`
}

// AddLicenseCmd ...
type AddLicenseCmd struct {
	Path      string `name:"path" help:"Paths to list." type:"path"`
	Extension string `name:"extension" help:"File extension you wish to licence."`
	License   string `help:"License file path to use."`
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	PluginXML  PluginXMLCmd  `cmd:"" help:"You can use the functions in that command to manipulate plugin.xml file under cordova plugin root project directory."`
	AddLicense AddLicenseCmd `cmd:"" help:"Add license to files."`
}
