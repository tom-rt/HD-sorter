package main

import (
	"os"
	"path/filepath"
)

var snap map[string]int = make(map[string]int)

func snapshot(path string) {
	items, _ := os.ReadDir(path)
	for _, item := range items {
		name := item.Name()
		if item.IsDir() && toSort(name, path) {
			snapshot(path + "/" + name)
		} else if !item.IsDir() && name != "hd-sorter" {
			ext := filepath.Ext(name)
			if ext == "" {
				snap["misc"] += 1
			} else {
				snap[ext] += 1
			}
		}
	}
}
