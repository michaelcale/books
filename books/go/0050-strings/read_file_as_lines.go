package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// :show start
// ReadFileAsLines reads a file and splits it into lines
func ReadFileAsLines(path string) ([]string, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := string(d)
	lines := strings.Split(s, "\n")
	return lines, nil
}

// :show end

func main() {
	path := "read_file_as_lines.go"
	lines, err := ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("ReadFileAsLines() failed with '%s'\n", err)
	}
	fmt.Printf("There are %d lines in '%s'\n", len(lines), path)
}
