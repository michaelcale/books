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

func genBookTOCJSON(book *Book) ([]byte, error) {
	bookTOC := BookTOC{
		Name: book.Title,
	}
	for _, ch := range book.Chapters {
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
		bookTOC.Chapters = append(bookTOC.Chapters, chtoc)
	}
	return json.MarshalIndent(&book, "", "  ")
}
