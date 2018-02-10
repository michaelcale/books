package main

import (
	"bufio"
	"os"
)

// :show start
// IterLinesInFile iterates over lines in a file
func IterLinesInFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Scan() reads next line and returns false when reached end or error
	for scanner.Scan() {
		line := scanner.Text()
		// process the line
	}
	// check if Scan() finished because of error or because it reached end of file
	return scanner.Err()
}

// :show end

func main() {
}
