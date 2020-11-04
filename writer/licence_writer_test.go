package writer

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCheckIfFileContainsLicenceAlready_LicenceExists(t *testing.T) {
	createFile("test.java")
	file, _ := os.OpenFile("test.java", os.O_RDWR, 0644)
	_, err := file.WriteString(readFile("licence"))
	_ = file.Close()
	ok, err := IsLicenceExist("test.java", "licence")
	if err != nil {
		t.Error()
	}
	if !ok {
		t.Error()
	}
	file.Close()
	os.Remove(file.Name())
}

func TestCheckIfFileContainsLicenceAlready_LicenceNotFound(t *testing.T) {
	createFile("test1.java")
	file, _ := os.OpenFile("test1.java", os.O_RDWR, 0644)
	ok, err := IsLicenceExist("test1.java", "licence")
	if err != nil {
		t.Error()
	}
	if ok {
		t.Error()
	}
	file.Close()
	os.Remove(file.Name())
}

func TestCheckIfFileContainsLicenceAlready_FileNotFound(t *testing.T) {
	_, err := IsLicenceExist("notFound.java", "licence")
	if err == nil {
		t.Error()
	}
}

func TestCheckIfFileContainsLicenceAlready_LicenceExistWithWrongFormat(t *testing.T) {
	createFile("test.java")
	file, _ := os.OpenFile("test.java", os.O_RDWR, 0644)
	licence := readFile("licence")
	lines := strings.Split(licence, "\n")
	lines[0] = "wrong formatter text"
	lines = append(lines, "wrong formatter text\n")
	licence = strings.Join(lines, "\n")
	if !isLicenceAlreadyContainsTags(licence, "/*") {
		licence = addTagToLicence("/*","*/",licence)
	}
	ioutil.WriteFile("testLicence", []byte(licence), 0644)
	_, _ = file.WriteString(licence)
	for i := 0; i < 30; i++ {
		file.WriteString("Test line\n")
	}
	file.Close()
	if IsLicenceFormatValid("test.java", "testLicence") {
		t.Error()
	}
	os.Remove(file.Name())
	os.Remove("testLicence")
}

func TestWriteToFileLicence_LicenceIsWrittenProperly(t *testing.T) {
	createFile("test.java")
	file, _ := os.OpenFile("test.java", os.O_RDWR, 0644)
	for i := 1; i < 100; i++ {
		_, _ = file.WriteString("This is a test file.\n")
	}
	_, _ = WriteLicenceToFile(file.Name(), "licence", "/*", "*/")
	s := readFile(file.Name())
	licence := readFile("licence")
	file.Close()
	os.Remove(file.Name())
	if !strings.Contains(s, licence) {
		t.Error()
	}
}

func createFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(""), 0644)
}
