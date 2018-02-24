package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// :show start
	d := []byte{0x01, 0xff, 0x3a, 0xcd}
	s := fmt.Sprintf("%x", d)
	fmt.Printf("Hex: %s\n", s)

	var decoded []byte
	_, err := fmt.Sscanf(s, "%x", &decoded)
	if err != nil {
		log.Fatalf("fmt.Sscanf() failed with '%s'\n", err)
	}
	if !bytes.Equal(d, decoded) {
		log.Fatalf("decoded version is different than original")
	}

	n := 3824
	fmt.Printf("%d in hex is 0x%x\n", n, n)
	// :show end
}
