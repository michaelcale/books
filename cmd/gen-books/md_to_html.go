package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdhtml "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/kjk/u"
	"github.com/microcosm-cc/bluemonday"
)

var (
	htmlFormatter  *html.Formatter
	highlightStyle *chroma.Style
)

func init() {
	htmlFormatter = html.New(html.WithClasses(), html.TabWidth(2))
	u.PanicIf(htmlFormatter == nil, "couldn't create html formatter")
	styleName := "monokailight"
	highlightStyle = styles.Get(styleName)
	u.PanicIf(highlightStyle == nil, "didn't find style '%s'", styleName)

}

// based on https://github.com/alecthomas/chroma/blob/master/quick/quick.go
func htmlHighlight(w io.Writer, source, lang, defaultLang string) error {
	if lang == "" {
		lang = defaultLang
	}
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return htmlFormatter.Format(w, highlightStyle, it)
}

func isArticleOrChapterLink(s string) bool {
	return strings.HasPrefix(s, "a-") || strings.HasPrefix(s, "ch-")
}

var didPrint = false

func printKnownURLS(a []string) {
	if didPrint {
		return
	}
	didPrint = true
	fmt.Printf("%d known urls\n", len(a))
	for _, s := range a {
		fmt.Printf("%s\n", s)
	}
}

// turn partial url like "a-20381" into a full url like "a-20381-installing"
func fixupURL(uri string, knownURLS []string) string {
	if !isArticleOrChapterLink(uri) {
		return uri
	}
	for _, known := range knownURLS {
		if uri == known {
			return uri
		}
		if strings.HasPrefix(known, uri) {
			//fmt.Printf("fixupURL: %s => %s\n", uri, known)
			return known
		}
	}
	fmt.Printf("fixupURL: didn't fix up: %s\n", uri)
	//printKnownURLS(knownURLS)
	return uri
}

// knownUrls is a list of chapter/article urls in the form "a-20381-installing", "ch-198-getting-started"
func makeRenderHookCodeBlock(defaultLang string, book *Book) mdhtml.RenderNodeFunc {
	return func(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
		if codeBlock, ok := node.(*ast.CodeBlock); ok {
			lang := string(codeBlock.Info)
			if false {
				fmt.Printf("lang: '%s', code: %s\n", lang, string(codeBlock.Literal[:16]))
				io.WriteString(w, "\n<pre class=\"chroma\"><code>")
				mdhtml.EscapeHTML(w, codeBlock.Literal)
				io.WriteString(w, "</code></pre>\n")
			} else {
				htmlHighlight(w, string(codeBlock.Literal), lang, defaultLang)
			}
			return ast.GoToNext, true
		} else if link, ok := node.(*ast.Link); ok {
			// we just want to fix up the url if it's just a prefix
			// of known url and let rendering to the original code
			dest := string(link.Destination)
			link.Destination = []byte(fixupURL(dest, book.knownUrls))
			return ast.GoToNext, false
		} else {
			return ast.GoToNext, false
		}
	}
}

func markdownToUnsafeHTML(md []byte, defaultLang string, book *Book) []byte {
	extensions := parser.NoIntraEmphasis |
		parser.Tables |
		parser.FencedCode |
		parser.Autolink |
		parser.Strikethrough |
		parser.SpaceHeadings |
		parser.NoEmptyLineBeforeBlock
	parser := parser.NewWithExtensions(extensions)

	htmlFlags := mdhtml.Smartypants |
		mdhtml.SmartypantsFractions |
		mdhtml.SmartypantsDashes |
		mdhtml.SmartypantsLatexDashes
	htmlOpts := mdhtml.RendererOptions{
		Flags:          htmlFlags,
		RenderNodeHook: makeRenderHookCodeBlock(defaultLang, book),
	}
	renderer := mdhtml.NewRenderer(htmlOpts)
	return markdown.ToHTML(md, parser, renderer)
}

func markdownToHTML(d []byte, defaultLang string, book *Book) string {
	unsafe := markdownToUnsafeHTML(d, defaultLang, book)
	policy := bluemonday.UGCPolicy()
	policy.AllowStyling()
	policy.RequireNoFollowOnFullyQualifiedLinks(false)
	policy.RequireNoFollowOnLinks(false)
	res := policy.SanitizeBytes(unsafe)
	return string(res)
}
