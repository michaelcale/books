package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/kjk/u"
)

func copyToWwwStaticMaybeMust(srcName string) {
	var dstPtr *string
	switch srcName {
	case "main.css":
		dstPtr = &pathMainCSS
	case "app.js":
		dstPtr = &pathAppJS
	case "font-awesome.min.js":
		dstPtr = &pathFontAwesomeJS
	default:
		u.PanicIf(true, "unknown srcName '%s'", srcName)
	}
	src := filepath.Join("tmpl", srcName)
	d, err := ioutil.ReadFile(src)
	u.PanicIfErr(err)
	sha1Hex := u.Sha1HexOfBytes(d)
	name := nameToSha1Name(srcName, sha1Hex)
	dst := filepath.Join("www", "s", name)
	err = ioutil.WriteFile(dst, d, 0644)
	u.PanicIfErr(err)
	*dstPtr = filepath.ToSlash(dst[len("www"):])
	fmt.Printf("Copied %s => %s\n", src, dst)
}

// data related to handling re-generation of book if source files change
// we respond to file system change notifications but want to debounce
// regeneration because they are expensive and operations like rename
// generate several notifications in a row
var (
	muRegen sync.Mutex
	// sequence number used debounce generation
	nextRegenSeq int32
	// books to regenerate since last regeneration
	booksToRegen map[string]struct{}
	// if true, regenerate all books
	regenAllBooks bool
)

// path is books/${book}/${chapter}/${article}
func getBookDirFromPath(path string) string {
	path = toUnixPath(path)
	if !strings.HasPrefix(path, "books/") {
		fmt.Printf("getBookDirFromPath('%s') => ''\n", path)
		return ""
	}
	path = strings.TrimPrefix(path, "books/")
	// now the path is "go/${chapter}/${article}.md"
	parts := strings.Split(path, "/")
	return parts[0]
}

func handleFileChange(path string) {
	fmt.Printf("handleFileChange: %s\n", path)

	name := filepath.Base(path)
	switch name {
	case "main.css", "app.js", "font-awesome.min.js":
		clearErrors()
		copyToWwwStaticMaybeMust(name)
		printAndClearErrors()
	}

	muRegen.Lock()
	defer muRegen.Unlock()

	nextRegenSeq++
	if strings.HasPrefix(path, "tmpl") {
		regenAllBooks = true
	} else {
		// we assume it's either .md file change or a directory rename
		book := getBookDirFromPath(path)
		if book != "" {
			if booksToRegen == nil {
				booksToRegen = make(map[string]struct{})
			}
			booksToRegen[book] = struct{}{}
		}
	}

	// wait a bit before regenerating books. this allows to collapse
	// multiple rapid changes into a single op
	go func(seq int32) {
		time.Sleep(time.Second * 3)
		muRegen.Lock()
		if seq != nextRegenSeq {
			// another file change arrived in the meantime, so we'll allow
			// next goroutine to re-generate changes
			muRegen.Unlock()
			return
		}

		var localBooksToRegen []string
		for book := range booksToRegen {
			localBooksToRegen = append(localBooksToRegen, book)
		}

		localRegenAllBooks := regenAllBooks

		regenAllBooks = false
		booksToRegen = nil
		muRegen.Unlock()

		clearErrors()
		unloadTemplates() // for reloading of templates from disk
		if localRegenAllBooks {
			genAllBooks()
		} else {
			genSelectedBooks(localBooksToRegen)
		}
		printAndClearErrors()
	}(nextRegenSeq)
}

// TODO: when a directory is renamed or created, I need to add it
// to the list of watched directories
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
