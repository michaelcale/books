package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func getDirsRecur(dir string) ([]string, error) {
	toVisit := []string{dir}
	idx := 0
	for idx < len(toVisit) {
		dir = toVisit[idx]
		idx++
		fileInfos, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		for _, fi := range fileInfos {
			if !fi.IsDir() {
				continue
			}
			path := filepath.Join(dir, fi.Name())
			toVisit = append(toVisit, path)
		}
	}
	return toVisit, nil
}
