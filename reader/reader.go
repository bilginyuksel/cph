package reader

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func getRootDir(path string) string {
	parts := strings.Split(path, string(os.PathSeparator))
	return parts[0]
}

// FilePathWalkDir ...
func FilePathWalkDir(root string, ignored []string) ([]string, error) {
	ignoredMap := make(map[string]bool)
	for _, value := range ignored {
		ignoredMap[value] = true
	}
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		rootDir := getRootDir(path)
		if _, ok := ignoredMap[rootDir]; ok {
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
