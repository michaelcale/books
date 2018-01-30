package main

import (
	"fmt"
	"log"

	"github.com/sony/sonyflake"
)

// Each chapter and article should have a unique id
// This program generates a random unique id that has a reasonably compact
// string representation

var (
	encodeTable = "0123456789abcdefghijklmnopqrstuv"
)

func encodeByte(b byte) (byte, byte) {
	n1 := b & 0x1f
	n2 := (b >> 4) & 0x1f
	return encodeTable[n1], encodeTable[n2]
}

func encode(n uint64) string {
	var res [8]byte
	for i := 0; i < 4; i++ {
		c := byte(n & 0xff)
		n = n >> 8
		c1, c2 := encodeByte(c)
		res[i*2] = c1
		res[i*2+1] = c2
	}
	return string(res[:])
}

func main() {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	//fmt.Printf("unique id: %x\n", id)
	fmt.Printf("unique id: %s\n", encode(id))
}
