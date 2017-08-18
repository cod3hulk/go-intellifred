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

var dirsToSkip = [...]string{"target",
	".git",
	".idea",
	"src",
	"bin", "lib",
	".mvn",
	".settings",
}

func skip(filename string) bool {
	for i := range dirsToSkip {
		if dirsToSkip[i] == filename {
			return true
		}
	}
	return false
}

func visit(path string, f os.FileInfo, err error) error {

	if f.IsDir() && skip(f.Name()) {
		return filepath.SkipDir
	}

	// search for idea config files
	if !f.IsDir() && filepath.Ext(f.Name()) == ".iml" {
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
	root := "/Volumes/development"
	filepath.Walk(root, visit)

	result := new(alfred.Result)
	result.AddAll(items)
	fmt.Print(result.Filter(os.Args[1], filterItem).Output())
}
