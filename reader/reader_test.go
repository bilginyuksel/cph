package reader

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func beforeTestFilePathFinderCreateFileStructure(t *testing.T) {
	os.Mkdir("test", 0755)
	os.Mkdir("test/test1", 0755)

	createEmptyFileForTesting := func(filename string) {
		d := []byte("")
		err := ioutil.WriteFile(filename, d, 0644)
		if err != nil {
			t.Error(err)
		}
	}
	createEmptyFileForTesting("test/test.go")
	createEmptyFileForTesting("test/test1/test1.cpp")
	createEmptyFileForTesting("test/test1/test1.go")
}

func afterTestFilePathFinderDestroyFileStructure(t *testing.T) {
	err := os.RemoveAll("test")
	if err != nil {
		t.Error(err)
	}
}

func TestFilePathFinder_ReturnFilesIncludedSubDirectories(t *testing.T) {
	beforeTestFilePathFinderCreateFileStructure(t)

	files, err := FilePathWalkDir("test", []string{})
	if err != nil {
		t.Errorf("Shouldn't error")
	}

	if len(files) < 0 {
		t.Errorf("Files length should be greater than 0")
	}

	expectedFilesWindows := [3]string{"test\\test.go", "test\\test1\\test1.cpp", "test\\test1\\test1.go"}
	expectedFilesLinux := [3]string{"test/test.go", "test/test1/test1.cpp", "test/test1/test1.go"}

	if len(expectedFilesWindows) != len(files) {
		t.Errorf("Expected files length= %d, Actual files length= %d", len(expectedFilesWindows), len(files))
	}

	fmt.Println(expectedFilesWindows)
	fmt.Println(expectedFilesLinux)
	actualFilesMap := make(map[string]bool)

	for i := 0; i < len(files); i++ {
		actualFilesMap[files[i]] = true
	}

	for i := 0; i < len(expectedFilesWindows); i++ {
		_, okWindows := actualFilesMap[expectedFilesWindows[i]]
		_, okLinux := actualFilesMap[expectedFilesLinux[i]]
		if !okWindows && !okLinux {
			t.Errorf("Expected files= %v, Actual files= %v", expectedFilesWindows, files)
		}
	}
	afterTestFilePathFinderDestroyFileStructure(t)
}

func TestFilePathFinder_ReturnFilesOnlySubdirectory(t *testing.T) {
	beforeTestFilePathFinderCreateFileStructure(t)
	files, err := FilePathWalkDir("test/test1", []string{})
	if err != nil {
		t.Error("Shouldn't error")
	}

	if len(files) < 0 {
		t.Errorf("Files length should be greater than 0")
	}

	expectedFilesWindows := [2]string{"test\\test1\\test1.cpp", "test\\test1\\test1.go"}
	expectedFilesLinux := [2]string{"test/test1/test1.cpp", "test/test1/test1.go"}

	if len(expectedFilesWindows) != len(files) {
		t.Errorf("Expected files= %v, Actual files= %v", expectedFilesWindows, files)
	}

	actualFilesMap := make(map[string]bool)

	for i := 0; i < len(files); i++ {
		actualFilesMap[files[i]] = true
	}

	for i := 0; i < len(expectedFilesWindows); i++ {
		_, okWindows := actualFilesMap[expectedFilesWindows[i]]
		_, okLinux := actualFilesMap[expectedFilesLinux[i]]
		if !okWindows && !okLinux {
			t.Errorf("File %s should be in found files. But actual files are %v", expectedFilesWindows[i], actualFilesMap)
		}
	}
	afterTestFilePathFinderDestroyFileStructure(t)
}

func TestFilePathFinder_NotFound(t *testing.T) {
	_, err := FilePathWalkDir("empty", []string{})
	if err == nil {
		t.Errorf("Expected error.")
	}
}
