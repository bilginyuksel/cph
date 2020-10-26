package writer

import (
	"fmt"
	"github.com/bilginyuksel/cordova-plugin-helper/parser"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCheckIfFileContainsLicenceAlready_LicenceExists(t *testing.T) {
	file, err := os.Create("test.java")
	_, err = file.WriteString(ReadFile("licence.java"))
	_ = file.Close()
	ok, err := CheckIfFileContainsLicenceAlready("test.java", "licence.java")
	if err != nil {
		t.Error()
	}
	if !ok {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestCheckIfFileContainsLicenceAlready_LicenceNotFound(t *testing.T) {
	file, err := os.Create("test1.java")
	_ = file.Close()
	ok, err := CheckIfFileContainsLicenceAlready("test1.java", "licence.java")
	if err != nil {
		t.Error()
	}
	if ok {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestCheckIfFileContainsLicenceAlready_FileNotFound(t *testing.T) {
	_, err := CheckIfFileContainsLicenceAlready("notFound.java", "licence.java")
	if err == nil {
		t.Error()
	}
}

func TestCheckIfFileContainsLicenceAlready_LicenceExistWithWrongFormat(t *testing.T) {
	file, err := os.Create("test.java")
	_, _ = file.WriteString(ReadFile("licence.java"))
	_, _ = file.WriteString("test string")
	_ = file.Close()
	file, err = os.OpenFile("licence.java", os.O_RDONLY, 0644)
	if err != nil {
		t.Error()
	}
	_ = file.Close()
	CheckIfLicenceFormatIsValid(file)
	os.Remove("test.java")
}

func TestWriteToFileLicence_LicenceIsWrittenProperly(t *testing.T) {
	file, err := os.Create("test.java")
	for i := 1; i < 100; i++ {
		file.WriteString("This is a test file.\n")
	}
	_ = file.Close()
	_, _ = WriteLicenceToFile(file.Name(), "licence.java")
	if err != nil {
		panic(err)
	}
	s := ReadFile(file.Name())
	if !strings.Contains(s, ReadFile("licence.java")) {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestWriteToFileLicence_LicenceIsNotWrittenProperly(t *testing.T) {
	ok, _ := WriteLicenceToFile("test.java", "licence.java")
	if ok {
		t.Error()
	}
}

func TestReadSourceFiles_AllJavaFilesAdded(t *testing.T) {
	os.Mkdir("src", 0755)
	os.Mkdir("www", 0755)
	file, err := os.Create("src/test1.java")
	plg, _ := parser.ParseXML("plugin.xml")
	parser.CreateXML(plg, "plg.xml")
	_, err = os.Stat("src/test1.java")
	if err != nil {
		if os.IsNotExist(err) {
			t.Error()
		}
	}
	file.Close()
	os.Remove(file.Name())
	os.RemoveAll("src")
	os.RemoveAll("www")
}

func TestReadJsModules_AllJsFilesAdded(t *testing.T) {
	os.Mkdir("src", 0755)
	os.Mkdir("www", 0755)
	file1, err := os.Create("www/test1.js")
	file2, err := os.Create("www/test2.js")
	file3, err := os.Create("www/test3.js")
	plg, _ := parser.ParseXML("plugin.xml")
	parser.CreateXML(plg, "plg.xml")
	for i := 1; i <= 3; i++ {
		_, err = os.Stat(fmt.Sprintf("www/test%d.js", i))
		if err != nil {
			if os.IsNotExist(err) {
				t.Error()
			}
		}
	}
	file1.Close()
	file2.Close()
	file3.Close()
	os.Remove(file1.Name())
	os.Remove(file2.Name())
	os.Remove(file3.Name())
	os.RemoveAll("src")
	os.RemoveAll("www")
}

func TestReadJsModules_PluginXmlDoesntContainNonExistFile(t *testing.T) {
	os.Mkdir("src", 0755)
	os.Mkdir("www", 0755)
	file1, _ := os.Create("www/test1.js")
	plg, _ := parser.ParseXML("plugin.xml")
	parser.CreateXML(plg, "plg.xml")
	file1.Close()
	os.Remove(file1.Name())
	resultPlugin, _ := parser.ParseXML("plg.xml")
	parser.CreateXML(resultPlugin, "plg.xml")
	b, _ := ioutil.ReadFile("plg.xml")
	content  := string(b)
	if strings.Contains(content,"www/test1.js") {
		t.Error()
	}
	os.Remove("plg.xml")
	os.RemoveAll("src")
	os.RemoveAll("www")
}
