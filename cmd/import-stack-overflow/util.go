package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// GzippedReadCloser is a io.ReadCloser for a gzip file
type GzippedReadCloser struct {
	f *os.File
	r io.Reader
}

// Close closes a reader
func (rc *GzippedReadCloser) Close() error {
	return rc.f.Close()
}

// Read reads data from a reader
func (rc *GzippedReadCloser) Read(d []byte) (int, error) {
	return rc.r.Read(d)
}

func openGzipped(path string) (io.ReadCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	rc := &GzippedReadCloser{
		f: f,
		r: r,
	}
	return rc, nil
}

func readGzipped(path string) ([]byte, error) {
	rc, err := openGzipped(path)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	d, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func jsonDecodeGzipped(path string, v interface{}) error {
	r, err := openGzipped(path)
	if err != nil {
		return err
	}
	defer r.Close()
	dec := json.NewDecoder(r)
	return dec.Decode(v)
}

// https://stackoverflow.com/questions/695438/safe-characters-for-friendly-url
func charIsURLSafe(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	switch c {
	case '-', '.', '_', '~':
		return true
	}
	return false
}

func shortenConsequitve(s string, c string) string {
	s2 := c + c
	for strings.Contains(s, s2) {
		s = strings.Replace(s, s2, c, -1)
	}
	// strip c from the beginning
	for len(s) > 0 && s[0:1] == c {
		s = s[1:]
	}
	return s
}

func makeURLSafe(s string) string {
	n := len(s)
	d := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		c := s[i]
		if charIsURLSafe(c) {
			if c == '.' {
				c = '-'
			}
			d = append(d, c)
		} else {
			if c == ' ' {
				d = append(d, '-')
			}
		}
	}
	s = string(d)
	s = strings.ToLower(s)
	s = shortenConsequitve(s, "-")
	return s
}
