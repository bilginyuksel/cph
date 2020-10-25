package writer

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCheckIfFileContainsLicenceAlready_LicenceExists(t *testing.T) {
	file, err := os.Create("test.java")
	_,err = file.WriteString(LICENCE)
	_ = file.Close()
	ok, err := CheckIfFileContainsLicenceAlready("test.java")
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
	ok, err := CheckIfFileContainsLicenceAlready("test1.java")
	if err != nil {
		t.Error()
	}
	if ok {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestCheckIfFileContainsLicenceAlready_FileNotFound(t *testing.T) {
	_, err := CheckIfFileContainsLicenceAlready("notFound.java")
	if err == nil {
		t.Error()
	}
}

func TestCheckIfFileContainsLicenceAlready_LicenceExistWithWrongFormat(t *testing.T) {
	file, err := os.Create("licence.java")
	_, _ = file.WriteString(LICENCE)
	_, _ = file.WriteString("test string")
	_ = file.Close()
	file, err = os.OpenFile("licence.java", os.O_RDONLY, 0644)
	if err != nil {
		t.Error()
	}
	_ = file.Close()
	CheckIfLicenceFormatIsValid(file)
	os.Remove(file.Name())
}

func TestWriteToFileLicence_LicenceIsWrittenProperly(t *testing.T) {
	file, err := os.Create("javaFileWithoutLicence.java")
	for i := 1; i<100; i++{
		file.WriteString("This is a test file.\n")
	}
	_ = file.Close()
	_, _ = WriteToFileLicence(file.Name())
	b, err := ioutil.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}
	s := string(b)
	if !strings.Contains(s, LICENCE) {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestWriteToFileLicence_LicenceIsNotWrittenProperly(t *testing.T) {
	ok, _ := WriteToFileLicence("test.java")
	if ok{
		t.Error()
	}
}
