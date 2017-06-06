package main

import (
	"fmt"
	"github.com/cod3hulk/alfred"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"os"
	"path/filepath"
	"strings"
)

var items []alfred.Item

func visit(path string, f os.FileInfo, err error) error {
	extension := filepath.Ext(f.Name())

	// ignore git config directories
	if f.IsDir() && extension == ".git" && extension != ".idea" {
		//fmt.Printf("Skipping .git dir")
		return filepath.SkipDir
	}

	// search for idea config files
	if extension == ".iml" {
		item := alfred.Item{
			Title:    strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())),
			Subtitle: path,
			Arg:      filepath.Dir(path),
		}
		items = append(items, item)
		return filepath.SkipDir
	}

	return nil
}

func filterItem(item alfred.Item, query string) bool {
	query = strings.ToLower(query)
	arg := strings.ToLower(item.Title)
	return fuzzy.Match(query, arg)
}

func main() {
	root := "/Users/tave/development"
	filepath.Walk(root, visit)

	result := new(alfred.Result)
	result.AddAll(items)
	fmt.Print(result.Filter(os.Args[1], filterItem).Output())
}
