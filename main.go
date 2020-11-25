package main

import (
	"fmt"

	lic "github.com/bilginyuksel/cph/licence"

	"github.com/bilginyuksel/cph/generator"
	"github.com/bilginyuksel/cph/parser"
	"github.com/bilginyuksel/cph/reader"

	"github.com/alecthomas/kong"
)

func main() {
	content := `/**
	* This is an interface.
	* @param value This is a random value.
	* @return Promise<boolean> This function returns anything.
	* @param callback callback function to pass bilmem ne.
	* nextnext
	*/
	function considerCase(value:number, callback: ()=>void = () => {console.log("hello world")}) {
	
	}`
	parser.Tokenize(content)
	parser.ParseLoop()
	//prepareCliParser()
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
	sourceFiles, _ := reader.FilePathWalkDir("src", []string{})
	plugin.Platform.NewSourceFrom(sourceFiles)
	jsModules, _ := reader.FilePathWalkDir("www", []string{})
	plugin.NewJsModulesFrom(jsModules)

	err = parser.CreateXML(plugin, "plugin.xml")
	return err
}

// AddLicenceTo ...
func AddLicenceTo(path string, ignored []string) error {
	fmt.Println(ignored)

	if path == "" {
		path = "."
	}
	files, _ := reader.FilePathWalkDir(path, ignored)

	for _, p := range files {
		lic.Write(p)
	}
	return nil
}

// PluginGenerator ...
func PluginGenerator(group string, project string, homePage string) error {
	generator.CreateBasePlugin(group, project, homePage)
	return nil
}

// Run ...
func (pl *PluginXMLCmd) Run(ctx *Context) error {
	return SyncPluginXML(pl.Path)
}

// Run ...
func (l *AddLicenseCmd) Run(ctx *Context) error {
	return AddLicenceTo(l.Path, l.Ignore)
}

// Run ...
func (p *PluginCmd) Run(ctx *Context) error {
	return PluginGenerator(p.Group, p.ProjectName, p.HomePage)
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
	Path   string   `name:"path" help:"Paths to list." type:"path" short:"p"`
	Ignore []string `name:"ignore" help:"list of ignored folders" short:"i"`
}

// PluginCmd ...
type PluginCmd struct {
	ProjectName string `required name:"project-name" help:"Project name for the plugin."`
	Group       string `required name:"domain" help:"Group name for the plugin."`
	HomePage    string `required name:"home-page" help:"Home page for the plugin."`
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Plugin     PluginCmd     `cmd:"" help:"Use this command to create a cordova plugin from scratch."`
	PluginXML  PluginXMLCmd  `cmd:"" help:"You can use the functions in that command to manipulate plugin.xml file under cordova plugin root project directory."`
	AddLicense AddLicenseCmd `cmd:"" help:"Add license to files."`
}
