package main

import (
	"fmt"
	"io"
	"os"
)

// :show start
// CopyFile copies a src file to dst
func CopyFile(dst, src string) error {
	fmt.Printf("CopyFile():\nsrc='%s'\ndst='%s'\n\n", src, dst)
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("os.Open('%s') failed with '%s'\n", src, err)
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Printf("os.Create('%s') failed with '%s'\n", dst, err)
		return err
	}
	_, err = io.Copy(dstFile, srcFile)
	err2 := dstFile.Close()
	if err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		// delete the destination if copy failed
		os.Remove(dst)
	}
	return err
}

// :show end

func main() {
	src := "file_copy.go"
	dst := "file_copy_copy.go"
	err := CopyFile(dst, src)
	if err != nil {
		fmt.Printf("CopyFile failed with '%s'\n", err)
	}
}
