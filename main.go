package main

import (
	"fmt"
	"github.com/cod3hulk/alfred"
	"os"
	"path/filepath"
	"strings"
)

var items []alfred.Item

func visit(path string, f os.FileInfo, err error) error {
	extension := filepath.Ext(f.Name())

	// ignore git config directories
	if f.IsDir() && extension == ".git" {
		//fmt.Printf("Skipping .git dir")
		return filepath.SkipDir
	}

	// search for idea config files
	if extension == ".iml" {
		item := alfred.Item{
			Title:    strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())),
			Subtitle: path,
			Arg:      path,
		}
		items = append(items, item)
		return filepath.SkipDir
	}

	return nil
}

func main() {
	root := "/Users/tave/development"
	filepath.Walk(root, visit)

	result := new(alfred.Result)
	result.AddAll(items)
	fmt.Printf(result.Output())
}
