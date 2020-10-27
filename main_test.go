package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bilginyuksel/cordova-plugin-helper/parser"
)

var (
	javaFiles []string = []string{"src/main/java/com/group/project/Test1.java",
		"src/main/java/com/group/project/Test2.java", "src/main/java/com/group/project/test/Test3.java",
		"src/main/java/com/group/project/test/test/Test4.java", "src/main/java/com/group/project/test/test/Test5.java",
		"src/main/java/com/group/project/test/test/test/Test6.java"}

	jsFiles []string = []string{"www/Test1.js", "www/test2.js", "www/test/test3.js",
		"www/test/Test4.js", "www/test/test/Test5.js", "www/test/test/test6.js"}
)

func createMockFileStructure() {
	os.Mkdir("src", 0755)
	os.Mkdir("src/main", 0755)
	os.Mkdir("src/main/java", 0755)
	os.Mkdir("src/main/java/com", 0755)
	os.Mkdir("src/main/java/com/group", 0755)
	os.Mkdir("src/main/java/com/group/project", 0755)
	os.Mkdir("src/main/java/com/group/project/test", 0755)
	os.Mkdir("src/main/java/com/group/project/test/test", 0755)
	os.Mkdir("src/main/java/com/group/project/test/test/test", 0755)

	os.Mkdir("www", 0755)
	os.Mkdir("www/test", 0755)
	os.Mkdir("www/test/test", 0755)

	for _, path := range javaFiles {
		ioutil.WriteFile(path, []byte(""), 0644)
	}

	for _, path := range jsFiles {
		ioutil.WriteFile(path, []byte(""), 0644)
	}

	ioutil.WriteFile("plugin.xml", []byte(`<?xml version='1.0' encoding='utf-8'?>
	<plugin id="cordova-plugin-hms-push"
			version="5.0.2"
			xmlns="http://apache.org/cordova/ns/plugins/1.0"
			xmlns:android="http://schemas.android.com/apk/res/android">
		<name>Cordova Plugin HMS Push</name>
		<description>Cordova Plugin HMS Push</description>
		<license>Apache 2.0</license>
		<keywords>android, huawei, hms, push</keywords>
	
		<engines>
			<engine name="cordova" version=">=3.0.0"/>
		</engines>

		<platform name="android">
		</platform>`), 0644)
}

func eraseMockFileStructure() {
	os.RemoveAll("src")
	os.RemoveAll("www")
	os.Remove("plugin.xml")
}
func TestSyncPluginXMLNoPathPluginXMLExists_UpdatePluginXML(t *testing.T) {
	createMockFileStructure()
	SyncPluginXML("")

	plugin, _ := parser.ParseXML("plugin.xml")

	if len(plugin.JsModule) != len(jsFiles) {
		t.Error()
	}

	if len(plugin.Platform.SourceFiles) != len(javaFiles) {
		t.Error()
	}

	jsFileMap := make(map[string]string)
	javaFileMap := make(map[string]string)

	for _, path := range jsFiles {
		_, name := filepath.Split(path)
		name = strings.TrimSuffix(name, filepath.Ext(path))
		jsFileMap[path] = name
	}

	for _, path := range javaFiles {
		dir, _ := filepath.Split(path)
		javaFileMap[path] = dir
	}

	fmt.Println(jsFileMap)
	_, ok := javaFileMap[strings.TrimSpace(plugin.Platform.SourceFiles[0].Src)]
	fmt.Println(ok)

	for i := 0; i < len(plugin.JsModule); i++ {
		currentJsModule := plugin.JsModule[i]
		name, isPresent := jsFileMap[currentJsModule.Src]
		if !isPresent || name != currentJsModule.Name {
			t.Errorf(`Current JS Module and the js module is not the same.
			Actual JS Module= %v
			Expected JS Module= %s`, currentJsModule, name)
		}
	}

	for i := 0; i < len(javaFiles); i++ {
		currentSourceFile := plugin.Platform.SourceFiles[i]
		dir, isPresent := javaFileMap[currentSourceFile.Src]
		if !isPresent || dir != currentSourceFile.TargetDir {
			t.Errorf(`Current Java Source File is not the same with the expected.
			Actual Source File= %v
			Expected Source File= %s`, currentSourceFile, javaFileMap[currentSourceFile.Src])
		}
	}

	eraseMockFileStructure()
}
