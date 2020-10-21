package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	resultFromHello := hello()

	if resultFromHello != "Hello, World" {
		t.Errorf("Expected result is 'Hello, World' but %s is given", resultFromHello)
	} else {
		t.Logf("Success.")
	}
}

func TestFilePathFinder_ReturnFilesIncludedSubDirectories(t *testing.T) {
	files, err := filePathWalkDir("test")
	if err != nil {
		t.Errorf("Shouldn't error")
	}

	if len(files) < 0 {
		t.Errorf("Files length should be greater than 0")
	}

	expectedFiles := [5]string{"test\\test.go", "test\\test.java",
		"test\\test1\\test1.cpp", "test\\test1\\test1.java", "test\\test1\\test1.go"}

	if len(expectedFiles) != len(files) {
		t.Errorf("Expected files length= %d, Actual files length= %d", len(expectedFiles), len(files))
	}

	fmt.Println(expectedFiles)
	actualFilesMap := make(map[string]bool)

	for i := 0; i < len(files); i++ {
		actualFilesMap[files[i]] = true
	}

	for i := 0; i < len(expectedFiles); i++ {
		_, ok := actualFilesMap[expectedFiles[i]]
		if !ok {
			t.Errorf("Expected files= %v, Actual files= %v", expectedFiles, files)
		}
	}
}

func TestFilePathFinder_ReturnFilesOnlySubdirectory(t *testing.T) {
	files, err := filePathWalkDir("test/test1")
	if err != nil {
		t.Error("Shouldn't error")
	}

	if len(files) < 0 {
		t.Errorf("Files length should be greater than 0")
	}

	expectedFiles := [3]string{"test\\test1\\test1.cpp", "test\\test1\\test1.java", "test\\test1\\test1.go"}

	if len(expectedFiles) != len(files) {
		t.Errorf("Expected files= %v, Actual files= %v", expectedFiles, files)
	}

	actualFilesMap := make(map[string]bool)

	for i := 0; i < len(files); i++ {
		actualFilesMap[files[i]] = true
	}

	for i := 0; i < len(expectedFiles); i++ {
		_, ok := actualFilesMap[expectedFiles[i]]
		if !ok {
			t.Errorf("File %s should be in found files. But actual files are %v", expectedFiles[i], actualFilesMap)
		}
	}

}

func TestFilePathFinder_NotFound(t *testing.T) {
	_, err := filePathWalkDir("empty")
	if err == nil {
		t.Errorf("Expected error.")
	}
}
