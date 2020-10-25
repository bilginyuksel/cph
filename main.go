package main

import (
	//"github.com/bilginyuksel/cordova-plugin-helper/writer"

	"os"
	"path/filepath"

	// "github.com/bilginyuksel/cordova-plugin-helper/parser"
	"github.com/alecthomas/kong"
)

func main() {
	// plg, _ := parser.ParseXML("parser/plugin.xml")
	// parser.CreateXML(plg, "plg.xml")
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return err
	})
	return files, err
}

// Context ...
type Context struct {
	Debug bool
}

// PluginXMLCmd ...
type PluginXMLCmd struct {
	Path string `help:"String to path"`
	Sync bool   `help:"Sync plugin.xml file. This command will search the related directories and if it finds any missing or unnecessary field it will add or delete automatically."`
}

// AddLicenseCmd ...
type AddLicenseCmd struct {
	Paths          []string `arg optional name:"path" help:"Paths to list." type:"path"`
	FileExtensions []string `arg optional name:"extension" help:"Instead of giving every files path, just give file extensions here."`
	License        string   `help:"License file path to use."`
}

// Run ...
// func (pl *PluginXMLCmd) Run(ctx *Context) error {
// 	fmt.Println(pl)
// 	return nil
// }

// Run ...
// func (l *AddLicenseCmd) Run(ctx *Context) error {
// 	fmt.Println("ls", l.Paths)
// 	return nil
// }

var cli struct {
	Debug bool `help:"Enable debug mode."`

	PluginXML  PluginXMLCmd  `cmd help:"You can use the functions in that command to manipulate plugin.xml file under cordova plugin root project directory."`
	AddLicense AddLicenseCmd `cmd help:"Add license to files."`
}
