package main

import (
	"fmt"
	"io"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gomarkdown/markdown"
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

func makeRenderHookCodeBlock(defaultLang string) markdown.RenderNodeFunc {
	return func(w io.Writer, node *markdown.Node, entering bool) (markdown.WalkStatus, bool) {
		nodeData, ok := node.Data.(*markdown.CodeBlockData)
		if !ok {
			return markdown.GoToNext, false
		}
		lang := string(nodeData.Info)
		if false {
			fmt.Printf("lang: '%s', code: %s\n", lang, string(node.Literal[:16]))
			io.WriteString(w, "\n<pre class=\"chroma\"><code>")
			markdown.EscapeHTML(w, node.Literal)
			io.WriteString(w, "</code></pre>\n")
		} else {
			htmlHighlight(w, string(node.Literal), lang, defaultLang)
		}
		return markdown.GoToNext, true
	}
}

func markdownToUnsafeHTML(md []byte, defaultLang string) []byte {
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
		Flags:          htmlFlags,
		RenderNodeHook: makeRenderHookCodeBlock(defaultLang),
	}
	renderer := markdown.NewHTMLRenderer(htmlParams)
	return markdown.ToHTML(md, parser, renderer)
}

func markdownToHTML(d []byte, defaultLang string) string {
	unsafe := markdownToUnsafeHTML(d, defaultLang)
	policy := bluemonday.UGCPolicy()
	policy.AllowStyling()
	policy.RequireNoFollowOnFullyQualifiedLinks(false)
	policy.RequireNoFollowOnLinks(false)
	res := policy.SanitizeBytes(unsafe)
	return string(res)
}
