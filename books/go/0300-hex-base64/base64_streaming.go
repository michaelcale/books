package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// :show start
	d := []byte{0x01, 0xff, 0x3a, 0xcd}

	writer := &bytes.Buffer{}
	base64Writer := base64.NewEncoder(base64.StdEncoding, writer)

	_, err := base64Writer.Write(d)
	if err != nil {
		log.Fatalf("base64Writer.Write() failed with '%s'\n", err)
	}
	err = base64Writer.Close()
	if err != nil {
		log.Fatalf("base64Writer.Close() failed with '%s'\n", err)
	}

	encoded := writer.Bytes()
	fmt.Printf("Base64: %s\n", string(encoded))

	reader := bytes.NewBuffer(encoded)
	base64Reader := base64.NewDecoder(base64.StdEncoding, reader)

	decoded, err := ioutil.ReadAll(base64Reader)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() failed with '%s'\n", err)
	}

	if !bytes.Equal(d, decoded) {
		log.Fatalf("decoded version is different than original")
	}
	// :show end
}
