package kvstore

import (
	"fmt"
	"io"
	"strings"

	"github.com/kjk/programming-books/pkg/common"
	"github.com/kjk/u"
)

const (
	// RecordSeparator is a (hopefully) unique string that separates records in Key/Value file
	RecordSeparator = "|======|"
)

// KeyValue represents a key/value pair
type KeyValue struct {
	Key   string
	Value string
}

// Doc is a series of KeyValue pairs
type Doc []KeyValue

// GetV finds value by key, returns an error if didn't find
func (d Doc) GetV(key string) (string, error) {
	for _, kv := range d {
		if kv.Key == key {
			return kv.Value, nil
		}
	}
	return "", fmt.Errorf("key '%s' not found", key)
}

// GetVSilent finds value by key. Returns def string if didn't find
func (d Doc) GetVSilent(key string, defValue string) string {
	for _, kv := range d {
		if kv.Key == key {
			return kv.Value
		}
	}
	return defValue
}

func extractMultiLineValue(lines []string) ([]string, string, error) {
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == RecordSeparator {
			rest := lines[i+1:]
			s := strings.Join(lines[:i], "\n")
			return rest, s, nil
		}
	}
	return nil, "", fmt.Errorf("didn't find end of value line ('%s')", RecordSeparator)
}

// if error is io.EOF, we successfully finished parsing
func parseNextKV(lines []string) ([]string, KeyValue, error) {
	// skip empty lines from the beginning
	var kv KeyValue
	for len(lines) > 0 && len(lines[0]) == 0 {
		lines = lines[1:]
	}
	if len(lines) == 0 {
		return nil, kv, io.EOF
	}
	s := strings.TrimSpace(lines[0])
	lines = lines[1:]

	if !strings.HasSuffix(s, ":") {
		// this is singlie line "k: v"
		parts := strings.SplitN(s, ":", 2)
		if len(parts) != 2 {
			return nil, kv, fmt.Errorf("'%s' is not a valid start for k/v", s)
		}
		kv.Key, kv.Value = parts[0], parts[1]
		return lines, kv, nil
	}
	// this is a multi-line value that ends with RecordSeparator
	kv.Key = strings.TrimSuffix(s, ":")
	var err error
	lines, kv.Value, err = extractMultiLineValue(lines)
	return lines, kv, err
}

/*
ParseKVFile parsers my brand of key/value text file optimized for human editing
Key/value are encoded in 2 ways:
1. On a single line, if value is short and doesn't contain '\n'

"key: value\n"

2. On multiple lines, if value is long or contains '\n'

key:
value
===\n
*/
func ParseKVFile(path string) (Doc, error) {
	lines, err := common.ReadFileAsLines(path)
	var res []KeyValue
	var kv KeyValue
	for {
		lines, kv, err = parseNextKV(lines)
		if err == io.EOF {
			return res, nil
		}
		if err != nil {
			return nil, err
		}
		res = append(res, kv)
	}
}

// can we serialize a given value on a single line or must use multiple lines?
func fitsOneLine(s string) bool {
	if len(s) > 80 {
		return false
	}
	if strings.Contains(s, "\n") {
		return false
	}
	// to avoid ambiguity when parsing serialize values with ':" on separate lines
	if strings.Contains(s, ":") {
		return false
	}
	return true
}

func isEmptyString(s string) bool {
	s = strings.TrimSpace(s)
	return len(s) == 0
}

// Serialize rserialized key/value
func Serialize(k, v string) string {
	if isEmptyString(v) {
		return ""
	}
	if fitsOneLine(v) {
		return fmt.Sprintf("%s: %s\n", k, v)
	}
	u.PanicIf(strings.Contains(v, RecordSeparator), "v contains RecordSeparator")

	return fmt.Sprintf("%s:\n%s\n%s\n", k, v, RecordSeparator)
}

// SerializeLong serializes key/value in the long form
func SerializeLong(k, v string) string {
	if isEmptyString(v) {
		return ""
	}
	u.PanicIf(strings.Contains(v, RecordSeparator), "v contains RecordSeparator")
	return fmt.Sprintf("%s:\n%s\n%s\n", k, v, RecordSeparator)
}
