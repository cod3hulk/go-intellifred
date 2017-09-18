package main

import (
	"flag"
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

func depth(path string) int {
	return len(strings.Split(path, "/"))
}

func visit(path string, f os.FileInfo, err error) error {

	if f.IsDir() && skip(f.Name()) {
		return filepath.SkipDir
	}

	if f.IsDir() && depth(path) > max_depth {
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

var root string
var project string
var max_depth int

func main() {
	flag.IntVar(&max_depth, "max-depth", 5, "Max directory depth to search for")
	flag.StringVar(&project, "project", "", "Project name to search for")
	flag.StringVar(&root, "root", "", "Directory where the projects are placed")
	flag.Parse()

	if project == "" {
		fmt.Println("Project ist mandatory")
		return
	}

	if root == "" {
		fmt.Println("Root ist mandatory")
		return
	}

	filepath.Walk(root, visit)

	result := new(alfred.Result)
	result.AddAll(items)
	fmt.Print(result.Filter(project, filterItem).Output())
}
