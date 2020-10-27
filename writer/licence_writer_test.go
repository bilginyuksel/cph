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
	_, err := file.WriteString(ReadFile("licence"))
	_ = file.Close()
	ok, err := CheckIfFileContainsLicenceAlready("test.java", "licence")
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
	ok, err := CheckIfFileContainsLicenceAlready("test1.java", "licence")
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
	_, err := CheckIfFileContainsLicenceAlready("notFound.java", "licence")
	if err == nil {
		t.Error()
	}
}

func TestCheckIfFileContainsLicenceAlready_LicenceExistWithWrongFormat(t *testing.T) {
	createFile("test.java")
	file, _ := os.OpenFile("test.java", os.O_RDWR, 0644)
	_, _ = file.WriteString(ReadFile("licence"))
	_, _ = file.WriteString("test string")
	file.Close()
	os.Remove(file.Name())
	file, err := os.OpenFile("licence", os.O_RDONLY, 0644)
	if err != nil {
		t.Error()
	}
	_ = file.Close()
	CheckIfLicenceFormatIsValid(file)
}

func TestWriteToFileLicence_LicenceIsWrittenProperly(t *testing.T) {
	createFile("test.java")
	file, _ := os.OpenFile("test.java", os.O_RDWR, 0644)
	for i := 1; i < 100; i++ {
		_, _ = file.WriteString("This is a test file.\n")
	}
	_, _ = WriteLicenceToFile(file.Name(), "licence")
	s := ReadFile(file.Name())
	licence := ReadFile("licence")
	file.Close()
	os.Remove(file.Name())
	if !strings.Contains(s, licence) {
		t.Error()
	}
}

func createFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(""), 0644)
}

