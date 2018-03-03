package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	os.MkdirAll("cached_output", 0755)
}

func cachedOutputPath(sha1Hex string) string {
	return filepath.Join("cached_output", sha1Hex)
}

// for a given file, get output of executing this command
// We cache this as it is the most expensive part of rebuilding books
// If allowError is true, we silence an error from executed command
// This is useful when e.g. executing "go run" on a program that is
// intentionally not valid.
func getCachedOutput(path string, allowError bool) (string, error) {
	fc, err := loadFileCached(path)
	if err != nil {
		return "", err
	}
	sha1Hex := fc.Sha1Hex()
	// fmt.Printf("getCachedOutput('%s'), sha1: %s\n", path, sha1Hex)
	outputPath := cachedOutputPath(sha1Hex)
	fc, err = loadFileCached(outputPath)
	if err == nil {
		return string(fc.Content), nil
	}
	// fmt.Printf("loadFileCached('%s') failed with '%s'\n", outputPath, err)
	s, err := getOutput(path)
	if err != nil {
		if !allowError {
			fmt.Printf("getOutput('%s'), output is:\n%s\n", path, s)
			return s, err
		}
		err = nil
	}
	d := []byte(s)
	outputPath = cachedOutputPath(sha1Hex)

	fmt.Printf("Wrote cached output '%s' for '%s'\n", outputPath, path)
	err = ioutil.WriteFile(outputPath, d, 0644)
	if err != nil {
		return "", err
	}
	return s, nil
}
