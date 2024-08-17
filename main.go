package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must pass 'sort' or 'snapshot' as an argument")
		return
	}

	mode := os.Args[1]
	if mode != "sort" && mode != "snapshot" {
		fmt.Println("you must pass 'sort' or 'snapshot' as an argument")
		return
	}

	if mode == "sort" {
		os.Mkdir("pictures", 0755)
		os.Mkdir("videos", 0755)
		os.Mkdir("documents", 0755)
		os.Mkdir("audios", 0755)
		sortFiles(".")
	} else if mode == "snapshot" {
		snapshot(".")
		// This block is to sort the files by occurences before printing result
		keys := make([]string, 0, len(snap))
		for key := range snap {
			keys = append(keys, key)
		}
		sort.SliceStable(keys, func(i, j int) bool {
			return snap[keys[i]] > snap[keys[j]]
		})
		for _, k := range keys {
			fmt.Println(k, snap[k])
		}
	}
}
