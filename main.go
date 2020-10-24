package main

import (
	"os"
	"path/filepath"

	parser "github.com/bilginyuksel/cordova-plugin-helper/parser"
)

func main() {
	plg, _ := parser.ParseXML("parser/plugin.xml")
	parser.CreateXML(plg, "plg.xml")
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
