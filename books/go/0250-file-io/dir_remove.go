package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	empty        = "empty_dir"
	nonEmpty     = "non_empty_dir"
	nonEmpty2Top = "non_empty_dir2"
	nonEmpty2Sub = filepath.Join(nonEmpty2Top, "second")
)

func createDir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatalf("os.MkdirAll('%s') failed with '%s'\n", path, err)
	}
}

func createFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("os.Create('%s') failed with '%s'\n", path, err)
	}
	err = f.Close()
	if err != nil {
		log.Fatalf("f.Close('%s') failed with '%s'\n", path, err)
	}
}

func createTestDirs() {
	os.RemoveAll(empty)
	os.RemoveAll(nonEmpty)
	os.RemoveAll(nonEmpty2Top)

	createDir(empty)
	createDir(nonEmpty)
	createDir(nonEmpty2Sub)

	path := filepath.Join(nonEmpty, "file.txt")
	createFile(path)
}

func remove(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Printf("os.Remove('%s') failed with '%s'\n", path, err)
	}
}

func removeAll(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Printf("os.RemoveAll('%s') failed with '%s'\n", path, err)
	}
}

func main() {
	createTestDirs()

	remove(empty)
	remove(nonEmpty)
	remove(nonEmpty2Top)

	removeAll(nonEmpty)
	removeAll(nonEmpty2Top)
}
