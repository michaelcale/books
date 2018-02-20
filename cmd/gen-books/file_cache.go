package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// FileContent describes a file and its content
type FileContent struct {
	Path    string
	Size    int64
	ModTime time.Time
	Content []byte
	Lines   []string
}

var (
	filePathToFileContent map[string]*FileContent
	muFileCache           sync.Mutex
)

func init() {
	filePathToFileContent = make(map[string]*FileContent)
}

func cacheFile(path string, info os.FileInfo) (*FileContent, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := string(d)
	lines := strings.Split(s, "\n")
	fc := &FileContent{
		Path:    path,
		Size:    info.Size(),
		ModTime: info.ModTime(),
		Content: d,
		Lines:   lines,
	}

	muFileCache.Lock()
	filePathToFileContent[path] = fc
	muFileCache.Unlock()

	return fc, nil
}

func cacheFileIfChanged(path string, info os.FileInfo) (*FileContent, error) {
	var err error
	if info == nil {
		info, err = os.Stat(path)
		if err != nil {
			return nil, err
		}
	}
	muFileCache.Lock()
	fc, ok := filePathToFileContent[path]
	muFileCache.Unlock()

	if !ok || fc.Size != info.Size() || fc.ModTime != info.ModTime() {
		return cacheFile(path, info)
	}
	return fc, nil
}

func cacheFilesCb(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	_, err = cacheFileIfChanged(path, info)
	return err
}

func cacheFilesInDir(dir string) error {
	return filepath.Walk(dir, cacheFilesCb)
}
