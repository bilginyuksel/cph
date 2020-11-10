package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/bilginyuksel/cph/reader"

	lic "github.com/bilginyuksel/cph/licence"
	"github.com/bilginyuksel/cph/parser"
)

var (
	linuxJavaFiles []string = []string{"src/main/java/com/group/project/Test1.java",
		"src/main/java/com/group/project/Test2.java", "src/main/java/com/group/project/test/Test3.java",
		"src/main/java/com/group/project/test/test/Test4.java", "src/main/java/com/group/project/test/test/Test5.java",
		"src/main/java/com/group/project/test/test/test/Test6.java"}

	winJavaFiles []string = []string{"src\\main\\java\\com\\group\\project\\Test1.java",
		"src\\main\\java\\com\\group\\project\\Test2.java", "src\\main\\java\\com\\group\\project\\test\\Test3.java",
		"src\\main\\java\\com\\group\\project\\test\\test\\Test4.java", "src\\main\\java\\com\\group\\project\\test\\test\\Test5.java",
		"src\\main\\java\\com\\group\\project\\test\\test\\test\\Test6.java"}

	linuxJsFiles []string = []string{"www/Test1.js", "www/test2.js", "www/test/test3.js",
		"www/test/Test4.js", "www/test/test/Test5.js", "www/test/test/test6.js"}

	winJsFiles []string = []string{"www\\Test1.js", "www\\test2.js", "www\\test\\test3.js",
		"www\\test\\Test4.js", "www\\test\\test\\Test5.js", "www\\test\\test\\test6.js"}
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

	for _, path := range linuxJavaFiles {
		ioutil.WriteFile(path, []byte(""), 0644)
	}

	for _, path := range linuxJsFiles {
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
		</platform>
		</plugin>`), 0644)
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

	if len(plugin.JsModule) != len(linuxJsFiles) {
		t.Error()
	}

	if len(plugin.Platform.SourceFiles) != len(linuxJavaFiles) {
		t.Error()
	}

	winJsFileMap := make(map[string]string)
	winJavaFileMap := make(map[string]string)
	linuxJsFileMap := make(map[string]string)
	linuxJavaFileMap := make(map[string]string)

	for _, path := range winJsFiles {
		_, name := filepath.Split(path)
		name = strings.TrimSuffix(name, filepath.Ext(path))
		winJsFileMap[path] = name
	}

	for _, path := range linuxJsFiles {
		_, name := filepath.Split(path)
		name = strings.TrimSuffix(name, filepath.Ext(path))
		linuxJsFileMap[path] = name
	}

	for _, path := range winJavaFiles {
		dir, _ := filepath.Split(path)
		winJavaFileMap[path] = dir
	}

	for _, path := range linuxJavaFiles {
		dir, _ := filepath.Split(path)
		linuxJavaFileMap[path] = dir
	}

	for i := 0; i < len(plugin.JsModule); i++ {
		currentJsModule := plugin.JsModule[i]
		ansWin, isPresentOnWin := winJsFileMap[currentJsModule.Src]
		ansLinux, isPresentOnLinux := linuxJsFileMap[currentJsModule.Src]
		if isPresentOnWin && ansWin == currentJsModule.Name {
			t.Logf("Passed windows type jsModule file paths... Data= %v", currentJsModule)
		} else if isPresentOnLinux && ansLinux == currentJsModule.Name {
			t.Logf("Passed linux type jsModule file paths... Data= %v", currentJsModule)
		} else {
			t.Error()
		}
	}

	for i := 0; i < len(linuxJavaFiles); i++ {
		currentSourceFile := plugin.Platform.SourceFiles[i]
		ansWin, isPresentOnWin := winJavaFileMap[currentSourceFile.Src]
		ansLinux, isPresentOnLinux := linuxJavaFileMap[currentSourceFile.Src]
		if isPresentOnWin && ansWin == currentSourceFile.TargetDir {
			t.Logf("Passed windows type sourceJava file paths... Data= %v", currentSourceFile)
		} else if isPresentOnLinux && ansLinux == currentSourceFile.TargetDir {
			t.Logf("Passed linux type sourceJava file paths... Data= %v", currentSourceFile)
		} else {
			t.Error()
		}
	}

	eraseMockFileStructure()
}

// func TestAddLicenceToJSFilesInCurrentPath_JSFilesShouldBeLicensed(t *testing.T) {
// 	createMockFileStructure()
// 	ioutil.WriteFile("www\\test.ts", []byte(""), 0644)
// 	winJsFiles = append(winJsFiles, "www\\test.ts")
// 	ioutil.WriteFile("www/test.ts", []byte(""), 0644)
// 	linuxJsFiles = append(linuxJsFiles, "www/test.ts")

// 	AddLicenceTo("www")
// 	if runtime.GOOS == "windows" {
// 		controlLicenceOnArrayOfFiles(winJsFiles, t, ".js")
// 	} else {
// 		controlLicenceOnArrayOfFiles(linuxJsFiles, t, ".js")
// 	}

// 	eraseMockFileStructure()
// }

func controlLicenceOnArrayOfFiles(files []string, t *testing.T, extension string) {
	for _, path := range files {
		ext := filepath.Ext(path)
		hasLicence := lic.IsExists(path)

		if hasLicence && ext == extension {
			t.Logf("Passed licence added to file= %s", path)
		} else if hasLicence && ext != extension {
			t.Errorf(`Failed there should be a licence if the file's extension is equal to '%s', 
			and there shouldn't be a licence if file's extension is different than '%s', File=%s`, extension, extension, path)
		} else if !hasLicence && ext == extension {
			t.Errorf(`Failed there should be a licence if the file's extension is equal to '%s', 
			and there shouldn't be a licence if file's extension is different than '%s', File=%s`, extension, extension, path)
		}
	}
}

// func TestAddLicenceToJavaFilesInCurrentPath_JavaFilesShouldBeLicensed(t *testing.T) {
// 	createMockFileStructure()
// 	ioutil.WriteFile("src\\test.ts", []byte(""), 0644)
// 	winJavaFiles = append(winJavaFiles, "src\\test.ts")
// 	ioutil.WriteFile("src/test.ts", []byte(""), 0644)
// 	linuxJavaFiles = append(linuxJavaFiles, "src/test.ts")

// 	AddLicenceTo("src")
// 	if runtime.GOOS == "windows" {
// 		controlLicenceOnArrayOfFiles(winJavaFiles, t, ".java")
// 	} else {
// 		controlLicenceOnArrayOfFiles(linuxJavaFiles, t, ".java")
// 	}

// 	eraseMockFileStructure()
// }

func TestAddLicenceToAllFilesInCurrentPath_AllFilesShouldBeLicensed(t *testing.T) {
	createMockFileStructure()
	ioutil.WriteFile("src\\test.ts", []byte(""), 0644)
	winJavaFiles = append(winJavaFiles, "src\\test.ts")
	ioutil.WriteFile("src/test.ts", []byte(""), 0644)
	linuxJavaFiles = append(linuxJavaFiles, "src/test.ts")

	checkAllFilesLicence := func(files []string) {
		for _, path := range files {
			content := reader.ReadFile(path)
			hasLicence := lic.IsExists(content)
			if !hasLicence {
				t.Error()
			}
		}
	}

	AddLicenceTo("src")
	if runtime.GOOS == "windows" {
		checkAllFilesLicence(winJavaFiles)
	} else {
		checkAllFilesLicence(linuxJavaFiles)
	}

	eraseMockFileStructure()
}
