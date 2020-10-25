package writer

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCheckIfFileContainsLicenceAlready_LicenceExists(t *testing.T) {
	ok, err := CheckIfFileContainsLicenceAlready("test.java")
	if err != nil {
		t.Error()
	}
	if !ok {
		t.Error()
	}
}

func TestCheckIfFileContainsLicenceAlready_LicenceNotFound(t *testing.T) {
	ok, err := CheckIfFileContainsLicenceAlready("test1.java")
	if err != nil {
		t.Error()
	}
	if ok {
		t.Error()
	}
}

func TestCheckIfFileContainsLicenceAlready_FileNotFound(t *testing.T) {
	_, err := CheckIfFileContainsLicenceAlready("notFound.java")
	if err == nil {
		t.Error()
	}
}

func TestCheckIfFileContainsLicenceAlready_LicenceExistWithWrongFormat(t *testing.T) {
	file, err := os.OpenFile("HmsPushMessaging.java", os.O_RDONLY, 0644)
	if err != nil {
		t.Error()
	}
	_ = file.Close()
	CheckIfLicenceFormatIsValid(file)
}

func TestWriteToFileLicence_LicenceIsWrittenProperly(t *testing.T) {
	_, _ = WriteToFileLicence("HmsPushMessaging.java")
	b, err := ioutil.ReadFile("HmsPushMessaging.java")
	if err != nil {
		panic(err)
	}
	s := string(b)
	if !strings.Contains(s, LICENCE) {
		t.Error()
	}
}

func TestWriteToFileLicence_LicenceIsNotWrittenProperly(t *testing.T) {
	ok, _ := WriteToFileLicence("test.java")
	if ok{
		t.Error()
	}
}
