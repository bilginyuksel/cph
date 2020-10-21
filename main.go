package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(hello())
	fmt.Println(filePathWalkDir("test"))
}

func hello() string {
	return "Hello, World"
}

func filePathWalkDir(root string) ([]string, error) {
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
