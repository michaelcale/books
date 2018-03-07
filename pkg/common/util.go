package common

import (
	"bytes"
	"strings"
)

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

// ShortenString shortens to max 60 characters
func ShortenString(s string) string {
	parts := strings.SplitN(s, "\n", 2)
	s = parts[0]
	if len(s) < 60 {
		return s
	}
	return s[:60] + "..."
}

// MakeURLSafe converts arbitrary string into a string that can be used as a file name or url
func MakeURLSafe(s string) string {

	// special consideration for "c#" etc.
	s = strings.Replace(s, "#", "sharp", -1)

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

// NormalizeNewlines normalizes \r\n (windows) and \r (mac)
// into \n (unix)
func NormalizeNewlines(d []byte) []byte {
	if len(d) == 0 {
		return d
	}
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return d
}
