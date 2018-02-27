package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	// :show start
	d := []byte{0x01, 0xff, 0x3a, 0xcd}
	s := base64.URLEncoding.EncodeToString(d)
	fmt.Printf("base64: %s\n", s)

	d2, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("hex.DecodeString() failed with '%s'\n", err)
	}
	if !bytes.Equal(d, d2) {
		log.Fatalf("decoded version is different than original")
	}
	// :show end
}
