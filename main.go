package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func visit(path string, f os.FileInfo, err error) error {
	extension := filepath.Ext(f.Name())

	// ignore git config directories
	if f.IsDir() && extension == ".git" {
		//fmt.Printf("Skipping .git dir")
		return filepath.SkipDir
	}

	// search for idea config files
	if extension == ".iml" {
		fmt.Printf("Visited: %s\n", f.Name())
		return filepath.SkipDir
	}

	return nil
}

func main() {
	root := "/Users/tave/development"
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
