package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// :show start
	d := []byte{0x01, 0xff, 0x3a, 0xcd}

	writer := &bytes.Buffer{}
	hexWriter := hex.NewEncoder(writer)

	_, err := hexWriter.Write(d)
	if err != nil {
		log.Fatalf("hexWriter.Write() failed with '%s'\n", err)
	}

	encoded := writer.Bytes()
	fmt.Printf("Hex: %s\n", string(encoded))

	reader := bytes.NewBuffer(encoded)
	hexReader := hex.NewDecoder(reader)

	decoded, err := ioutil.ReadAll(hexReader)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() failed with '%s'\n", err)
	}

	if !bytes.Equal(d, decoded) {
		log.Fatalf("decoded version is different than original")
	}
	// :show end
}
