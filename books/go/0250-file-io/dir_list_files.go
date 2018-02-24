package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// :show start
func main() {
	dir := filepath.Join("books", "go", "0250-file-io")
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("ioutil.ReadDir('%s') failed with '%s'\n", dir, err)
	}
	for i, fi := range fileInfos {
		if i < 4 {
			fmt.Printf("Path: %s, is dir: %v, size: %d bytes\n", fi.Name(), fi.IsDir(), fi.Size())
		}
	}
}

// :show end
