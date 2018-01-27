package main

import (
	"bytes"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

var (
	useChroma = true
)

func markdownToUnsafeHTML(md []byte) []byte {
	extensions := markdown.NoIntraEmphasis |
		markdown.Tables |
		markdown.FencedCode |
		markdown.Autolink |
		markdown.Strikethrough |
		markdown.SpaceHeadings |
		markdown.NoEmptyLineBeforeBlock
	parser := markdown.NewParserWithExtensions(extensions)

	htmlFlags := markdown.Smartypants |
		markdown.SmartypantsFractions |
		markdown.SmartypantsDashes |
		markdown.SmartypantsLatexDashes
	htmlParams := markdown.HTMLRendererParameters{
		Flags: htmlFlags,
	}
	renderer := markdown.NewHTMLRenderer(htmlParams)
	return markdown.ToHTML(md, parser, renderer)
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
