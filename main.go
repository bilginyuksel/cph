package main

import (
	"fmt"
	"path/filepath"
	"strings"

	lic "github.com/bilginyuksel/cph/licence"

	"github.com/bilginyuksel/cph/generator"
	"github.com/bilginyuksel/cph/parser"
	"github.com/bilginyuksel/cph/reader"
	"github.com/bilginyuksel/cph/tsc"

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
	sourceFiles, _ := reader.FilePathWalkDir("src", []string{})
	plugin.Platform.NewSourceFrom(sourceFiles)
	jsModules, _ := reader.FilePathWalkDir("www", []string{})
	plugin.NewJsModulesFrom(jsModules)
	if len(plugin.Platform.ConfigFiles) == 0 {
		splittedPluginIDString := strings.Split(plugin.ID, "-")
		pluginName := "Example"
		if len(splittedPluginIDString) > 1 {
			pluginName = splittedPluginIDString[len(splittedPluginIDString)-1]
		}
		firstLetterUpper := strings.ToUpper(string(pluginName[0])) + pluginName[1:]
		configFile := parser.ConfigFile{Target: "config.xml", Parent: "/*", Features: []parser.Feature{parser.Feature{Name: "HMS" + firstLetterUpper, Params: []parser.Param{parser.Param{Name: "android-package", Value: fmt.Sprintf("com.huawei.hms.cordova.%s.HMS%s", pluginName, firstLetterUpper)}}}}}
		plugin.Platform.ConfigFiles = append(plugin.Platform.ConfigFiles, configFile)
	}

	// ADD HOOKS
	if len(plugin.Platform.Hooks) == 0 {
		hooks := parser.Hook{Src: "hooks/before_plugin_uninstall.js", Type: "before_plugin_uninstall"}
		plugin.Platform.Hooks = append(plugin.Platform.Hooks, hooks)
		hooks = parser.Hook{Src: "hooks/after_plugin_install.js", Type: "after_plugin_install"}
		plugin.Platform.Hooks = append(plugin.Platform.Hooks, hooks)
		hooks = parser.Hook{Src: "hooks/after_prepare.js", Type: "after_prepare"}
		plugin.Platform.Hooks = append(plugin.Platform.Hooks, hooks)
	}
	// ADD HOOKS

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
func PluginGenerator(project string, include bool, tsutils bool) error {
	if include {
		generator.IncludeFramework(project)
	} else if tsutils {
		generator.CreateTSUtil()
	} else {
		generator.CreateHMSPlugin(project)
	}
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
	return PluginGenerator(p.Project, p.Include, p.TSUtils)
}

func getJavaFileNames(files []string) []string {
	javaFiles := []string{}
	for _, val := range files {
		ext := filepath.Ext(val)
		if ext == ".java" {
			javaFiles = append(javaFiles, val)
		}
	}
	return javaFiles
}

func getAllCormetReferences(javaFiles []string) []tsc.CormetRef {
	corRefList := []tsc.CormetRef{}
	for _, val := range javaFiles {
		_, value := filepath.Split(val)
		value = strings.Replace(value, ".java", "", -1)
		content := reader.ReadFile(val)
		if tsc.HasCormet(content) {
			corRefList = append(corRefList, *tsc.GetCormetRef(content, value))
		}
	}
	return corRefList
}

// Run ...
func (g *GenerateCmd) Run(ctx *Context) error {

	if g.TypeScript {
		files, _ := reader.FilePathWalkDir(".", []string{})
		javaFiles := getJavaFileNames(files)
		cormetRefList := getAllCormetReferences(javaFiles)
		if g.Single && len(g.Filename) > 0 {
			tsc.WriteCormetRefListToFile(g.Filename, cormetRefList)
		} else {
			tsc.WriteCormetRefListToFiles(cormetRefList)
		}
	}
	return nil
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
	Create  bool   `group:"choice" xor:"choice"`
	Include bool   `group:"choice" xor:"choice"`
	TSUtils bool   `name:"ts-utils" help:"Creates the base ts utils file." short:"ti"`
	Project string `name:"project" help:"Project name for the plugin."`
}

// GenerateCmd ...
type GenerateCmd struct {
	TypeScript bool   `required:"" name:"typescript" short:"t"`
	Single     bool   `name:"single" short:"s" help:"Create all references inside one file instead of seperate files."`
	Filename   string `name:"filename" short:"f" help:"To use single command you have to write the filename."`
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Plugin     PluginCmd     `cmd:"" help:"Use this command to create a cordova plugin from scratch."`
	PluginXML  PluginXMLCmd  `cmd:"" help:"You can use the functions in that command to manipulate plugin.xml file under cordova plugin root project directory."`
	Generate   GenerateCmd   `cmd:"" help:"Auto generation tool."`
	AddLicense AddLicenseCmd `cmd:"" help:"Add license to files."`
}
