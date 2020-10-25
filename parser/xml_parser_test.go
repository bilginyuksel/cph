package parser

import (
	"testing"
)

func TestParseFile_ReturnCorrectFileInformation(t *testing.T) {
	xmlResult, err := ParseXML("plugin.xml")
	if err != nil {
		t.Error()
	}

	if xmlResult.ID != "cordova-plugin-hms-push" && xmlResult.Author != "" &&
		xmlResult.License != "Apache 2.0" && xmlResult.Description != "Cordova Plugin HMS Push" &&
		xmlResult.Name != "Cordova Plugin HMS Push" {
		t.Logf("Actual: %s, Expected: %s", xmlResult.ID, "cordova-plugin-hms-push")
		t.Logf("Actual: %s, Expected: %s", xmlResult.Author, "")
		t.Logf("Actual: %s, Expected: %s", xmlResult.License, "Apache 2.0")
		t.Logf("Actual: %s, Expected: %s", xmlResult.Description, "Cordova Plugin HMS Push")
		t.Logf("Actual: %s, Expected: %s", xmlResult.Name, "Cordova Plugin HMS Push")
		t.Error()
	}

}

func TestParseFile_FileNotFound(t *testing.T) {
	_, err := ParseXML("notfound.xml")
	if err == nil {
		t.Error()
	}
}

func TestParseFile_ExtensionError(t *testing.T) {
	_, err := ParseXML("file.yaml")
	if err == nil {
		t.Error()
	}
}
