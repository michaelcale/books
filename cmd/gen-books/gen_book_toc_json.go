package main

import "encoding/json"

// BookSectionTOC represents a section in a book TOC
type BookSectionTOC struct {
	Title string `json:"title"`
	// TODO: url?
}

// BookChapterTOC represents chapter in book TOC
type BookChapterTOC struct {
	Title    string           `json:"title"`
	Sections []BookSectionTOC `json:"sections"`
}

// BookTOC represents table of contents of a book
type BookTOC struct {
	Name     string           `json:"name"`
	Chapters []BookChapterTOC `json:"chapters"`
}

func genBookTOCJSON(bookName string, chapters []*Chapter) ([]byte, error) {
	book := BookTOC{
		Name: bookName,
	}
	for _, ch := range chapters {
		var sections []BookSectionTOC
		for _, sec := range ch.Sections {
			sectoc := BookSectionTOC{
				Title: sec.Title,
			}
			sections = append(sections, sectoc)
		}
		chtoc := BookChapterTOC{
			Title:    ch.Title,
			Sections: sections,
		}
		book.Chapters = append(book.Chapters, chtoc)
	}
	return json.MarshalIndent(&book, "", "  ")
}
