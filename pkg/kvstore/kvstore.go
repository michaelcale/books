package kvstore

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/essentialbooks/books/pkg/common"
	"github.com/kjk/u"
)

const (
	// recordSeparator is a (hopefully) unique string that separates records in Key/Value file
	recordSeparator = "|======|"
)

// KeyValue represents a key/value pair
type KeyValue struct {
	Key   string
	Value string
}

// Doc is a series of KeyValue pairs
type Doc []KeyValue

// Get finds value by key, returns an error if didn't find
func (d Doc) Get(key string) (string, error) {
	for _, kv := range d {
		if kv.Key == key {
			return kv.Value, nil
		}
	}
	return "", fmt.Errorf("key '%s' not found", key)
}

// GetSilent finds value by key. Returns def string if didn't find
func (d Doc) GetSilent(key string, defValue string) string {
	for _, kv := range d {
		if kv.Key == key {
			return kv.Value
		}
	}
	return defValue
}

// ReplaceOrAppend appends key/value to doc and returns a potentially new doc
func ReplaceOrAppend(doc Doc, key, value string) Doc {
	for idx := range doc {
		kv := &doc[idx]
		if kv.Key == key {
			kv.Value = value
			return doc
		}
	}
	el := KeyValue{
		Key:   key,
		Value: value,
	}
	return append(doc, el)
}

func extractMultiLineValue(lines []string) ([]string, string, error) {
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == recordSeparator {
			rest := lines[i+1:]
			s := strings.Join(lines[:i], "\n")
			return rest, s, nil
		}
	}
	// for convenience, we don't requuire RecordSeparator because most
	// articles have only one multi-line value and it ends at the end of filea
	// This might not detect mistakenly leaving RecordSeparator when we have
	// more than one multi-line value. We can live with that.
	s := strings.Join(lines[:], "\n")
	return nil, s, nil
	//return nil, "", fmt.Errorf("didn't find end of value line ('%s')", RecordSeparator)
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
		parts := strings.SplitN(s, ": ", 2)
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

func isYamlSeparator(s string) bool {
	s = strings.TrimSpace(s)
	return s == "---"
}

// parses a variant of the file which has yaml metadata at the top
// https://github.com/blog/1647-viewing-yaml-metadata-in-your-documents
// The body is represented as Body key
func parseKVFileWithYamlMeta(lines []string) (Doc, error) {
	line := lines[0]
	lines = lines[1:]
	u.PanicIf(!isYamlSeparator(line), "first line is '%s' and should be '---'", line)
	var res []KeyValue
	var kv KeyValue
	var err error
	for len(lines) > 0 {
		if isYamlSeparator(lines[0]) {
			body := strings.Join(lines[1:], "\n")
			kv = KeyValue{
				Key:   "Body",
				Value: body,
			}
			res = append(res, kv)
			return res, nil
		}
		lines, kv, err = parseNextKV(lines)
		if err != nil {
			return nil, err
		}
		res = append(res, kv)
	}
	return nil, errors.New("didn't find closing '---'")
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
	if err != nil {
		return nil, err
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("%s is an empty document", path)
	}
	return ParseKVLines(lines)
}

// ParseKVLines parses KV format from lines
func ParseKVLines(lines []string) (Doc, error) {
	if isYamlSeparator(lines[0]) {
		return parseKVFileWithYamlMeta(lines)
	}
	var (
		res []KeyValue
		kv  KeyValue
		err error
	)
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
	u.PanicIf(strings.Contains(v, recordSeparator), "v contains RecordSeparator")

	return fmt.Sprintf("%s:\n%s\n%s\n", k, v, recordSeparator)
}

// SerializeLong serializes key/value in the long form
func SerializeLong(k, v string) string {
	if isEmptyString(v) {
		return ""
	}
	u.PanicIf(strings.Contains(v, recordSeparator), "v contains RecordSeparator")
	return fmt.Sprintf("%s:\n%s\n%s\n", k, v, recordSeparator)
}

// SerializeDoc serializes the doc in the new format where
// header contains all metadata information and the rest is Body
/*
---
Id: 3
Title: My title
---
Body
*/
func SerializeDoc(doc Doc) (string, error) {
	var lines = []string{"---"}
	var body string
	hasBody := false
	for _, kv := range doc {
		if kv.Key == "Body" {
			body = kv.Value
			hasBody = true
			continue
		}
		v := strings.TrimSpace(kv.Value)
		if strings.Contains(v, "\n") {
			return "", fmt.Errorf("key: '%s' value '%s' contains \\n", kv.Key, v)
		}
		if len(v) > 256 {
			return "", fmt.Errorf("key: '%s', value is %d bytes (> 256)", kv.Key, v)
		}
		s := fmt.Sprintf("%s: %s", kv.Key, v)
		lines = append(lines, s)
	}
	if !hasBody {
		id := doc.GetSilent("Id", "")
		return "", fmt.Errorf("Doc with id '%s' has no body", id)
	}
	lines = append(lines, "---")
	lines = append(lines, "") // for readability
	lines = append(lines, body)
	// TODO: remove duplicate empty lines
	return strings.Join(lines, "\n"), nil
}
