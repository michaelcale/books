package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	// :show start
	d := []byte{0x01, 0xff, 0x3a, 0xcd}
	s := hex.EncodeToString(d)
	fmt.Printf("Hex: %s\n", s)

	d2, err := hex.DecodeString(s)
	if err != nil {
		log.Fatalf("hex.DecodeString() failed with '%s'\n", err)
	}
	if !bytes.Equal(d, d2) {
		log.Fatalf("decoded version is different than original")
	}
	// :show end
}
