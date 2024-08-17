package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func sortFiles(path string) {
	items, _ := os.ReadDir(path)
	for _, item := range items {
		name := item.Name()
		if item.IsDir() && toSort(name, path) {
			sortFiles(path + "/" + name)
		} else if !item.IsDir() && name != "hd-sorter" {
			ext := filepath.Ext(name)
			if checkType(ext, pictureExt) {
				moveFile(path+"/"+name, "pictures/"+name)
			} else if checkType(ext, videoExt) {
				moveFile(path+"/"+name, "videos/"+name)
			} else if checkType(ext, docExt) {
				moveFile(path+"/"+name, "documents/"+name)
			} else if checkType(ext, audioExt) {
				moveFile(path+"/"+name, "audios/"+name)
			}
		}
	}
}

func moveFile(oldLocation string, newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		fmt.Println("an error occured while moving", oldLocation)
	}

}
