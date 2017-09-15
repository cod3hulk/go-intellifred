package main

import (
	"fmt"
	"github.com/cod3hulk/alfred"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"os"
	"path/filepath"
	"strconv"
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

func depth(path string) int {
	return len(strings.Split(path, "/"))
}

func visit(path string, f os.FileInfo, err error) error {

	if f.IsDir() && skip(f.Name()) {
		return filepath.SkipDir
	}

	if f.IsDir() && depth(path) > Max_depth {
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

var args = os.Args[1:]
var Root = args[0]
var Filename = args[1]
var Max_depth, err = strconv.Atoi(args[2])

func main() {
	if err != nil {
		Max_depth = 5
	}

	filepath.Walk(Root, visit)

	result := new(alfred.Result)
	result.AddAll(items)
	fmt.Print(result.Filter(Filename, filterItem).Output())
}
