package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// :show start

func main() {
	path := filepath.Join("books", "go", "0250-file-io", "file_info.go")
	st, err := os.Stat(path)
	if err != nil {
		log.Fatalf("GetFileSize failed with '%s'\n", err)
	}
	fmt.Printf(`Name: %s
Size: %d
IsDir: %v
Mode: %x
ModTime: %s
OS info: %T
`, st.Name(), st.Size(), st.IsDir(), st.Mode, st.ModTime(), st.Sys())
}

// :show end
