package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/kjk/programming-books/pkg/common"
	"github.com/kjk/programming-books/pkg/kvstore"
	"github.com/kjk/programming-books/pkg/stackoverflow"
	"github.com/kjk/u"
)

type DocTag = stackoverflow.DocTag
type Topic = stackoverflow.Topic
type Example = stackoverflow.Example
type TopicHistory = stackoverflow.TopicHistory
type Contributor = stackoverflow.Contributor

var (
	gDocTags        []DocTag
	gTopics         []Topic
	gExamples       []*Example
	gTopicHistories []TopicHistory
	gContributors   []*Contributor
	currDefaultLang string

	emptyExamplexs []*Example
	// if true, prints more information
	verbose = false

	booksToImport = common.BooksToProcess
)

func mdToHTML(d []byte) []byte {
	return markdown.ToHTML(d, nil, nil)
}

func getTopicsByDocID(docID int) map[int]bool {
	topics := make(map[int]bool)
	for _, topic := range gTopics {
		if topic.DocTagId == docID {
			topics[topic.Id] = true
		}
	}
	return topics
}

func isEmptyString(s string) bool {
	s = strings.TrimSpace(s)
	return len(s) == 0
}

func calcExampleCount(docTag *DocTag) {
	docID := docTag.Id
	topics := getTopicsByDocID(docID)
	n := 0
	for _, ex := range gExamples {
		if topics[ex.DocTopicId] {
			n++
		}
	}
	docTag.ExampleCount = n
}

func printDocTagsMust() {
	loadAll()
	for i := range gDocTags {
		docTag := &gDocTags[i]
		calcExampleCount(docTag)
	}
	sort.Slice(gDocTags, func(i, j int) bool {
		return gDocTags[i].ExampleCount < gDocTags[j].ExampleCount
	})
	for _, dc := range gDocTags {
		fmt.Printf(`{ "%s", "", false, %d, %d },%s`, dc.Title, dc.ExampleCount, dc.TopicCount, "\n")
	}
}

func loadDocTagsMust() []DocTag {
	path := path.Join("stack-overflow-docs-dump", "doctags.json.gz")
	docTags, err := stackoverflow.LoadDocTags(path)
	u.PanicIfErr(err)
	return docTags
}

func loadTopicsMust() []Topic {
	path := path.Join("stack-overflow-docs-dump", "topics.json.gz")
	topics, err := stackoverflow.LoadTopics(path)
	u.PanicIfErr(err)
	return topics
}

func loadTopicHistoriesMust() []TopicHistory {
	path := path.Join("stack-overflow-docs-dump", "topichistories.json.gz")
	topicHistories, err := stackoverflow.LoadTopicHistories(path)
	u.PanicIfErr(err)
	return topicHistories
}

func loadContributorsMust() []*Contributor {
	path := path.Join("stack-overflow-docs-dump", "contributors.json.gz")
	contributors, err := stackoverflow.LoadContibutors(path)
	u.PanicIfErr(err)
	return contributors
}

func loadExamplesMust() []*Example {
	path := path.Join("stack-overflow-docs-dump", "examples.json.gz")
	examples, err := stackoverflow.LoadExamples(path)
	u.PanicIfErr(err)
	return examples
}

func findDocTagByTitleMust(docTags []DocTag, title string) DocTag {
	for _, dc := range docTags {
		if dc.Title == title {
			return dc
		}
	}
	log.Fatalf("Didn't find DocTag with title '%s'\n", title)
	return DocTag{}
}

func loadAll() {
	timeStart := time.Now()
	fmt.Printf("Loading Stack Overflow data...")
	gDocTags = loadDocTagsMust()
	gTopics = loadTopicsMust()
	gExamples = loadExamplesMust()
	gTopicHistories = loadTopicHistoriesMust()
	gContributors = loadContributorsMust()
	fmt.Printf(" took %s\n", time.Since(timeStart))
}

func getTopicsByDocTagID(docTagID int) []*Topic {
	var res []*Topic
	for i, topic := range gTopics {
		if topic.DocTagId == docTagID {
			res = append(res, &gTopics[i])
		}
	}
	return res
}

func getExampleByID(id int) *Example {
	for i, e := range gExamples {
		if e.Id == id {
			return gExamples[i]
		}
	}
	return nil
}

func getExamplesForTopic(docTagID int, docTopicID int) []*Example {
	var res []*Example
	seenIds := make(map[int]bool)
	for _, th := range gTopicHistories {
		if th.DocTagId == docTagID && th.DocTopicId == docTopicID {
			id := th.DocExampleId
			if seenIds[id] {
				continue
			}
			seenIds[id] = true
			ex := getExampleByID(id)
			if ex == nil {
				//fmt.Printf("Didn't find example, docTagID: %d, docTopicID: %d\n", docTagID, docTopicID)
			} else {
				res = append(res, ex)
			}
		}
	}
	return res
}

func sortExamples(a []*Example) {
	sort.Slice(a, func(i, j int) bool {
		if a[i].IsPinned {
			return true
		}
		if a[j].IsPinned {
			return false
		}
		return a[i].Score > a[j].Score
	})
}

func shortenVersion(s string) string {
	if s == "[]" {
		return ""
	}
	return s
}

func writeIndexTxtMust(path string, topic *Topic) {
	s := kvstore.Serialize("Title", topic.Title)
	s += kvstore.Serialize("Id", strconv.Itoa(topic.Id))
	versions := shortenVersion(topic.VersionsJson)
	s += kvstore.SerializeLong("Versions", versions)
	if isEmptyString(versions) {
		s += kvstore.SerializeLong("VersionsHtml", topic.HelloWorldVersionsHtml)
	}

	s += kvstore.SerializeLong("Introduction", topic.IntroductionMarkdown)
	if isEmptyString(topic.IntroductionMarkdown) {
		s += kvstore.SerializeLong("IntroductionHtml", topic.IntroductionHtml)
	}

	s += kvstore.SerializeLong("Syntax", topic.SyntaxMarkdown)
	if isEmptyString(topic.SyntaxMarkdown) {
		s += kvstore.SerializeLong("SyntaxHtml", topic.SyntaxHtml)
	}

	s += kvstore.SerializeLong("Parameters", topic.ParametersMarkdown)
	if isEmptyString(topic.ParametersMarkdown) {
		s += kvstore.SerializeLong("ParametersHtml", topic.ParametersHtml)
	}

	s += kvstore.SerializeLong("Remarks", topic.RemarksMarkdown)
	if isEmptyString(topic.RemarksMarkdown) {
		s += kvstore.SerializeLong("RemarksHtml", topic.RemarksHtml)
	}

	createDirForFileMust(path)
	err := ioutil.WriteFile(path, []byte(s), 0644)
	u.PanicIfErr(err)
	if verbose {
		fmt.Printf("Wrote %s, %d bytes\n", path, len(s))
	}
}

func writeArticleMust(path string, example *Example) {
	s := kvstore.Serialize("Title", example.Title)
	s += kvstore.Serialize("Id", strconv.Itoa(example.Id))
	s += kvstore.Serialize("Score", strconv.Itoa(example.Score))
	s += kvstore.SerializeLong("Body", example.BodyMarkdown)
	if isEmptyString(example.BodyMarkdown) {
		s += kvstore.SerializeLong("BodyHtml", example.BodyHtml)
	}

	createDirForFileMust(path)
	err := ioutil.WriteFile(path, []byte(s), 0644)
	u.PanicIfErr(err)
	if verbose {
		fmt.Printf("Wrote %s, %d bytes\n", path, len(s))
	}
}

func printEmptyExamples() {
	for _, ex := range emptyExamplexs {
		fmt.Printf("empty example: %s, len(BodyHtml): %d\n", ex.Title, len(ex.BodyHtml))
	}
}

func getContributors(docID int) []int {
	topics := getTopicsByDocID(docID)
	contributors := make(map[int]bool)
	for _, c := range gContributors {
		topicID := c.DocTopicId
		if _, ok := topics[topicID]; ok {
			contributors[c.UserId] = true
		}
	}
	var res []int
	for id := range contributors {
		res = append(res, id)
	}
	return res
}

func genContributors(bookDstDir string, docID int) {
	contributors := getContributors(docID)
	var a []string
	for _, id := range contributors {
		a = append(a, strconv.Itoa(id))
	}
	s := strings.Join(a, "\n")
	path := filepath.Join(bookDstDir, "so_contributors.txt")
	createDirForFileMust(path)
	err := ioutil.WriteFile(path, []byte(s), 0644)
	u.PanicIfErr(err)
	//fmt.Printf("Wrote %s\n", path)
}

// this generates human-readable TOC, just for easy human inspection of the book structure
// potentially could be used to re-order book content without renaming files, just
// moving around
func genTOCTxtMust(path string, docID int) {
	var lines []string
	topics := getTopicsByDocTagID(docID)
	for _, t := range topics {
		s := fmt.Sprintf("%s %d", t.Title, t.Id)
		lines = append(lines, s)
		examples := getExamplesForTopic(docID, t.Id)
		sortExamples(examples)
		for _, ex := range examples {
			s := fmt.Sprintf("  %s %d", ex.Title, ex.Id)
			lines = append(lines, s)
		}
	}
	s := strings.Join(lines, "\n")
	err := ioutil.WriteFile(path, []byte(s), 0644)
	u.PanicIfErr(err)
}

func genBook(book *common.Book, defaultLang string) {
	timeStart := time.Now()
	name := book.Name
	newName := book.NewName()
	currDefaultLang = defaultLang
	bookDstDir := common.MakeURLSafe(newName)
	docTag := findDocTagByTitleMust(gDocTags, name)
	//fmt.Printf("%s: docID: %d\n", title, docTag.Id)
	topics := getTopicsByDocTagID(docTag.Id)
	nChapters := len(topics)
	nArticles := 0
	chapter := 10
	for _, t := range topics {
		examples := getExamplesForTopic(docTag.Id, t.Id)
		sortExamples(examples)

		dirChapter := fmt.Sprintf("%04d-%s", chapter, common.MakeURLSafe(t.Title))
		dirPath := filepath.Join("books", bookDstDir, dirChapter)
		chapterIndexPath := filepath.Join(dirPath, "index.md")
		writeIndexTxtMust(chapterIndexPath, t)
		//fmt.Printf("%s\n", dirChapter)
		chapter += 10
		//fmt.Printf("%s, %d examples (%d), %s\n", t.Title, t.ExampleCount, len(examples), fileName)

		articleNo := 10
		for _, ex := range examples {
			if isEmptyString(ex.BodyMarkdown) && isEmptyString(ex.BodyHtml) {
				emptyExamplexs = append(emptyExamplexs, ex)
				continue
			}
			fileName := fmt.Sprintf("%03d-%s.md", articleNo, common.MakeURLSafe(ex.Title))
			path := filepath.Join(dirPath, fileName)
			writeArticleMust(path, ex)
			//fmt.Printf("  %s %s '%s'\n", ex.Title, pinnedStr, fileName)
			//fmt.Printf("  %03d-%s\n", articleNo, fileName)
			//fmt.Printf("  %s\n", fileName)
			articleNo += 10
		}
		nArticles += len(examples)
	}
	bookDstPath := filepath.Join("books", bookDstDir)
	genContributors(bookDstPath, docTag.Id)
	path := filepath.Join(bookDstPath, "toc.txt")
	genTOCTxtMust(path, docTag.Id)

	fmt.Printf("Imported %s (%d chapters, %d articles) in %s\n", name, nChapters, nArticles, time.Since(timeStart))
}

func main() {
	if false {
		printDocTagsMust()
		return
	}
	timeStart := time.Now()
	loadAll()
	for _, bookInfo := range booksToImport {
		if !bookInfo.Import {
			continue
		}
		genBook(bookInfo, "")
	}
	fmt.Printf("Took %s\n", time.Since(timeStart))
	printEmptyExamples()
}
