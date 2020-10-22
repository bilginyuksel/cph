package main

import (
	"github.com/bilginyuksel/cordova-plugin-helper/writer"
	"os"
	"path/filepath"
	// parser "github.com/bilginyuksel/cordova-plugin-helper/parser"
)

func main() {
	// parser.Run()
	files, _ := filePathWalkDir("test/test2")
	writer.Run(files)
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
