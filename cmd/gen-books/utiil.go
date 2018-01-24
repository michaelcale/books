package main

import (
	"bytes"
)

func normalizeNewlines(d []byte) []byte {
	// replace CR LF (windows) with LF (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF (mac) with LF (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return d
}

// return first line of d and the rest
func bytesRemoveFirstLine(d []byte) (string, []byte) {
	idx := bytes.IndexByte(d, 10)
	//u.PanicIf(-1 == idx)
	if -1 == idx {
		return string(d), nil
	}
	l := d[:idx]
	return string(l), d[idx+1:]
}
