package main

import "testing"

func TestHello(t *testing.T) {
	resultFromHello := hello()

	if resultFromHello != "Hello, World" {
		t.Errorf("Expected result is 'Hello, World' but %s is given", resultFromHello)
	} else {
		t.Logf("Success.")
	}
}

func TestFilePathFinder_ShouldReturnFiles(t *testing.T) {
	files, err := filePathWalkDir("test/")
	if err != nil {
		t.Errorf("Shouldn't error")
	}

	if len(files) < 0 {
		t.Errorf("Files length should be greater than 0")
	}
}

func TestFilePathFinder_NotFound(t *testing.T) {
	_, err := filePathWalkDir("empty/")
	if err == nil {
		t.Errorf("Expected error.")
	}
}
