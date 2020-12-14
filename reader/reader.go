package reader

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func isIgnored(path string, ignored []string) bool {
	ignoredMap := make(map[string]bool)
	for _, value := range ignored {
		ignoredMap[value] = true
	}
	folders := strings.Split(path, "\\")
	for _, folder := range folders {
		if _, ok := ignoredMap[folder]; ok {
			return true
		}
	}
	return false
}

// FilePathWalkDir ...
func FilePathWalkDir(root string, ignored []string) ([]string, error) {

	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if isIgnored(path, ignored) {
			return nil
		}
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return err
	})
	return files, err
}

// ReadFile ...
func ReadFile(filename string) string {
	bytes, _ := ioutil.ReadFile(filename)
	return string(bytes)
}
