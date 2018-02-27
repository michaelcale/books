package main

import (
	"fmt"
	"path/filepath"
	"sync"
)

// SoContributor describes a StackOverflow contributor
type SoContributor struct {
	ID      int
	URLPart string
	Name    string
}

// Book represents a book
type Book struct {
	Title          string // used in index.tmpl.html
	titleSafe      string
	TitleLong      string // used in book_index.tmpl.html
	FileNameBase   string
	Chapters       []*Chapter
	sourceDir      string // dir where source markdown files are
	destDir        string // dif where destitation html files are
	SoContributors []SoContributor

	cachedArticlesCount int
	defaultLang         string // default programming language for programming examples
	knownUrls           []string

	// for concurrency
	sem chan bool
	wg  sync.WaitGroup
}

// ContributorCount returns number of contributors
func (b *Book) ContributorCount() int {
	return len(b.SoContributors)
}

// ContributorsURL returns url of the chapter that lists contributors
func (b *Book) ContributorsURL() string {
	return b.URL() + "/a-contributors"
}

// GitHubText returns text we show in GitHub link
func (b *Book) GitHubText() string {
	return "Edit on GitHub"
}

// GitHubURL returns link to GitHub for this book
func (b *Book) GitHubURL() string {
	return gitHubBaseURL + "/tree/master/books/" + filepath.Base(b.destDir)
}

// URL returns url of the book, used in index.tmpl.html
func (b *Book) URL() string {
	return fmt.Sprintf("/essential/%s/", b.titleSafe)
}

// CanonnicalURL returns full url including host
func (b *Book) CanonnicalURL() string {
	return urlJoin(siteBaseURL, b.URL())
}

// ShareOnTwitterText returns text for sharing on twitter
func (b *Book) ShareOnTwitterText() string {
	return fmt.Sprintf(`"Essential %s" - a free programming book`, b.Title)
}

// TocSearchJSURL returns data for searching titles of chapters/articles
func (b *Book) TocSearchJSURL() string {
	return b.URL() + "toc_search.js"
}

// CoverURL returns url to cover image
func (b *Book) CoverURL() string {
	coverName := langToCover[b.titleSafe]
	return fmt.Sprintf("/covers/%s.png", coverName)
}

// CoverFullURL returns a URL for the cover including host
func (b *Book) CoverFullURL() string {
	return urlJoin(siteBaseURL, b.CoverURL())
}

// CoverTwitterFullURL returns a URL for the cover including host
func (b *Book) CoverTwitterFullURL() string {
	coverName := langToCover[b.titleSafe]
	coverURL := fmt.Sprintf("/covers/twitter/%s.png", coverName)
	return urlJoin(siteBaseURL, coverURL)
}

// ArticlesCount returns total number of articles
func (b *Book) ArticlesCount() int {
	if b.cachedArticlesCount != 0 {
		return b.cachedArticlesCount
	}
	n := 0
	for _, ch := range b.Chapters {
		n += len(ch.Articles)
	}
	// each chapter has 000-index.md which is also an article
	n += len(b.Chapters)
	b.cachedArticlesCount = n
	return n
}

// ChaptersCount returns number of chapters
func (b *Book) ChaptersCount() int {
	return len(b.Chapters)
}
