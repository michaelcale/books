package main

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
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
