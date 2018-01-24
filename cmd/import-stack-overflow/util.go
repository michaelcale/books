package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kjk/u"
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

func createDirForFileMust(path string) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	u.PanicIfErr(err)
}
