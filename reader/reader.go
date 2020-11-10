package reader

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// FilePathWalkDir ...
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
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
