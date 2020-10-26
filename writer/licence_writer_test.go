package writer

import (
	"os"
	"strings"
	"testing"
)

func TestCheckIfFileContainsLicenceAlready_LicenceExists(t *testing.T) {

	file, err := os.Create("test.java")
	_, err = file.WriteString(ReadFile("licence"))
	_ = file.Close()
	ok, err := CheckIfFileContainsLicenceAlready("test.java", "licence")
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
	ok, err := CheckIfFileContainsLicenceAlready("test1.java", "licence")
	if err != nil {
		t.Error()
	}
	if ok {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestCheckIfFileContainsLicenceAlready_FileNotFound(t *testing.T) {
	_, err := CheckIfFileContainsLicenceAlready("notFound.java", "licence")
	if err == nil {
		t.Error()
	}
}

func TestCheckIfFileContainsLicenceAlready_LicenceExistWithWrongFormat(t *testing.T) {
	file, err := os.Create("test.java")
	_, _ = file.WriteString(ReadFile("licence"))
	_, _ = file.WriteString("test string")
	_ = file.Close()
	file, err = os.OpenFile("licence", os.O_RDONLY, 0644)
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
	_, _ = WriteLicenceToFile(file.Name(), "licence")
	if err != nil {
		panic(err)
	}
	s := ReadFile(file.Name())
	if !strings.Contains(s, ReadFile("licence")) {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestWriteToFileLicence_LicenceIsNotWrittenProperly(t *testing.T) {
	ok, _ := WriteLicenceToFile("test.java", "licence")
	if ok {
		t.Error()
	}
}
