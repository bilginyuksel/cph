package main

import (
	"fmt"

	"github.com/bilginyuksel/cordova-plugin-helper/parser"
	"github.com/bilginyuksel/cordova-plugin-helper/reader"

	"github.com/alecthomas/kong"
	//"github.com/bilginyuksel/cordova-plugin-helper/writer"
)

func main() {
	prepareCliParser()
}

func prepareCliParser() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
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
	Paths          []string `name:"path" help:"Paths to list." type:"path"`
	FileExtensions []string `name:"extension" help:"Instead of giving every files path, just give file extensions here."`
	License        string   `help:"License file path to use."`
}

// Run ...
func (pl *PluginXMLCmd) Run(ctx *Context) error {
	if pl.Path == "" {
		pl.Path = "."
	}
	files, _ := reader.FilePathWalkDir(pl.Path)
	fmt.Println(files)
	plg, _ := parser.ParseXML(fmt.Sprintf("%s/plugin.xml", pl.Path))
	fmt.Println(plg)
	// javaFiles, _ := reader.FilePathWalkDir("src")
	// plg.Platform.NewSourceFrom(javaFiles)
	// jsModules, _ := reader.FilePathWalkDir("www")
	// plg.NewJsModulesFrom(jsModules)
	// parser.CreateXML(plg, "plugin.xml")
	// fmt.Println(pl)
	return nil
}

// Run ...
// func (l *AddLicenseCmd) Run(ctx *Context) error {
// 	fmt.Println("ls", l.Paths)
// 	return nil
// }

var cli struct {
	Debug bool `help:"Enable debug mode."`

	PluginXML  PluginXMLCmd  `cmd:"" help:"You can use the functions in that command to manipulate plugin.xml file under cordova plugin root project directory."`
	AddLicense AddLicenseCmd `cmd:"" help:"Add license to files."`
}
