package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/kjk/u"
)

var (
	softErrorMode bool
	errors        []string
)

func maybePanicIfErr(err error) {
	if err == nil {
		return
	}
	if !softErrorMode {
		u.PanicIfErr(err)
	}
	errors = append(errors, err.Error())
}

func clearErrors() {
	errors = nil
}

func printAndClearErrors() {
	if len(errors) == 0 {
		return
	}
	errStr := strings.Join(errors, "\n")
	fmt.Printf("\n%d errors:\n%s\n\n", len(errors), errStr)
	clearErrors()
}

func createDirForFileMaybeMust(path string) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	maybePanicIfErr(err)
}

func copyFileMaybeMust(dst, src string) {
	createDirForFileMaybeMust(dst)
	err := copyFile(dst, src)
	maybePanicIfErr(err)
}

func copyToWwwMaybeMust(path string) {
	name := filepath.Base(path)
	dst := filepath.Join("www", name)
	copyFileMaybeMust(dst, path)
	fmt.Printf("Copied %s => %s\n", path, dst)
}

var (
	muRegen       sync.Mutex
	nextRegenSeq  int32
	booksToRegen  []string
	regenAllBooks bool
)

func handleFileChange(path string) {
	fmt.Printf("handleFileChange: %s\n", path)

	// those happen fast so we can just do them
	if strings.HasSuffix(path, "main.css") {
		clearErrors()
		copyToWwwMaybeMust(filepath.Join("tmpl", "main.css"))
		printAndClearErrors()
		return
	}

	if strings.HasSuffix(path, "app.js") {
		clearErrors()
		copyToWwwMaybeMust(filepath.Join("tmpl", "app.js"))
		printAndClearErrors()
		return
	}

	muRegen.Lock()
	defer muRegen.Unlock()

	nextRegenSeq++
	if strings.HasSuffix(path, ".tmpl.html") {
		regenAllBooks = true
	} else if strings.HasSuffix(path, ".md") {
		// TODO: figure out which book is it and add to booksToRegen
	} else {
		// most likely a directory moved
		// TODO: figure out which book is it and add to booksToRegen
	}

	// wait a second before regenerating books. this allows to collapse
	// multiple rapid changes into a single op
	go func(seq int32) {
		time.Sleep(time.Second)
		muRegen.Lock()
		if seq != nextRegenSeq {
			// another file change arrived in the meantime, so we'll allow
			// next goroutine to re-generate changes
			muRegen.Unlock()
			return
		}

		//localRegenAllBooks := regenAllBooks
		regenAllBooks = false
		//localBooksToRegen := booksToRegen
		booksToRegen = nil
		muRegen.Unlock()

		clearErrors()
		// TODO: use localRegenAllBooks and localBooksToRegen
		unloadTemplates() // for reloading of templates from disk
		genAllBooks()
		printAndClearErrors()
	}(nextRegenSeq)
}

func rebuildOnChanges() {
	softErrorMode = true
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

	<-done
	fmt.Printf("exited rebuildOnChanges()\n")
	os.Exit(1)
}
