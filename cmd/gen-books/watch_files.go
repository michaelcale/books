package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/kjk/u"
)

func copyToWwwMust(path string) {
	name := filepath.Base(path)
	dst := filepath.Join("www", name)
	copyFileMust(dst, path)
}

func handleFileChange(path string) {
	fmt.Printf("handleFileChange: %s\n", path)

	if strings.HasSuffix(path, "main.css") {
		copyToWwwMust(filepath.Join("tmpl", "main.css"))
		return
	}

	if strings.HasSuffix(path, "app.js") {
		copyToWwwMust(filepath.Join("tmpl", "app.js"))
		return
	}

	if strings.HasSuffix(path, ".tmpl.html") {
		fmt.Printf("Template changed, rebuilding all books\n")
		unloadTemplates() // for reloading of templates from disk
		genAllBooks()
		return
	}

	if strings.HasSuffix(path, ".md") {
		fmt.Printf("Rebuilding all books\n")
		// TODO: only rebuild the article or just the book
		// TODO: this doesn't pick up new files
		genAllBooks()
		return
	}

	// if this is rename of a directory, the name is the old name, so the directory
	// no longer exists
	// assume this is a renamed chapter directory
	// TODO: only rebuild the book that changed
	fmt.Printf("Rebuilding all books\n")
	genAllBooks()
}

func rebuildOnChanges() {
	dirs, err := getDirsRecur("tmpl")
	u.PanicIfErr(err)
	dirs2, err := getDirsRecur("books")
	u.PanicIfErr(err)
	dirs = append(dirs, dirs2...)

	watcher, err := fsnotify.NewWatcher()
	u.PanicIfErr(err)
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered in rebuildOnChanges(). Error: '%s'\n", r)
				done <- true
			}
		}()

		for {
			select {
			case event := <-watcher.Events:
				// filter out events that are just chmods
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					continue
				}
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
				}
				handleFileChange(event.Name)
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()
	for _, dir := range dirs {
		//fmt.Printf("Watching dir: '%s'\n", dir)
		watcher.Add(dir)
	}
	// waiting forever
	// TODO: pick up ctrl-c and cleanup and quit
	<-done
	fmt.Printf("exited rebuildOnChanges()\n")
	os.Exit(1)
}
