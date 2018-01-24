package main

import (
	"bytes"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var (
	useChroma = true
)

func markdownToUnsafeHTML(text []byte) []byte {
	// Those are blackfriday.MarkdownCommon() extensions
	/*
		extensions := 0 |
			EXTENSION_NO_INTRA_EMPHASIS |
			EXTENSION_TABLES |
			EXTENSION_FENCED_CODE |
			EXTENSION_AUTOLINK |
			EXTENSION_STRIKETHROUGH |
			EXTENSION_SPACE_HEADERS |
			EXTENSION_HEADER_IDS |
			EXTENSION_BACKSLASH_LINE_BREAK |
			EXTENSION_DEFINITION_LISTS
	*/

	// https://github.com/shurcooL/github_flavored_markdown/blob/master/main.go#L82
	extensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK

	commonHTMLFlags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	renderer := blackfriday.HtmlRenderer(commonHTMLFlags, "", "")
	opts := blackfriday.Options{Extensions: extensions}
	return blackfriday.MarkdownOptions(text, renderer, opts)
}

func markdownToHTML(s []byte) string {
	var replacements map[string][]byte
	if useChroma {
		s, replacements = markdownCodeHighligh(s)
	} else {
		s, replacements = txtWithCodeParts(s)
	}

	unsafe := markdownToUnsafeHTML(s)
	//unsafe := blackfriday.MarkdownCommon(s)
	policy := bluemonday.UGCPolicy()
	policy.AllowStyling()
	policy.RequireNoFollowOnFullyQualifiedLinks(false)
	policy.RequireNoFollowOnLinks(false)
	res := policy.SanitizeBytes(unsafe)

	// restore code snippets
	for kstr, v := range replacements {
		k := []byte(kstr)
		res = bytes.Replace(res, k, v, -1)
	}
	return string(res)
}
