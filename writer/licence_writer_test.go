package writer

import (
	"os"
	"testing"
)

func TestCheckIfLicenceExists_LicenceExists(t *testing.T) {
	ok, err := CheckIfLicenceExists("test.java")
	if err != nil {
		t.Error()
	}
	if !ok {
		t.Error()
	}
}

func TestCheckIfLicenceExists_LicenceNotFound(t *testing.T) {
	ok, err := CheckIfLicenceExists("test1.java")
	if err != nil {
		t.Error()
	}
	if ok {
		t.Error()
	}
}

func TestCheckIfLicenceExists_FileNotFound(t *testing.T) {
	_, err := CheckIfLicenceExists("notFound.java")
	if err == nil {
		t.Error()
	}
}

func TestCheckIfLicenceExists_LicenceExistWithWrongFormat(t *testing.T) {
	file, err := os.OpenFile("HmsPushMessaging.java",os.O_RDONLY,0644)
	if err != nil {
		t.Error()
	}
	CheckIfLicenceFormatIsValid(file)
}
