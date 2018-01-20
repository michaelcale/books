package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/kjk/u"
)

var (
	gDocTags        []DocTag
	gTopics         []Topic
	gExamples       []Example
	gTopicHistories []TopicHistory
)

func printDocTagsMust() {
	docTags := loadDocTagsMust()
	sort.Slice(docTags, func(i, j int) bool {
		return docTags[i].TopicCount < docTags[j].TopicCount
	})
	for _, dc := range docTags {
		fmt.Printf("%s: %s, %d\n", dc.Title, dc.Tag, dc.TopicCount)
	}
}

func loadDocTagsMust() []DocTag {
	path := path.Join("stack-overflow-docs-dump", "doctags.json.gz")
	docTags, err := loadDocTags(path)
	u.PanicIfErr(err)
	return docTags
}

func loadTopicsMust() []Topic {
	path := path.Join("stack-overflow-docs-dump", "topics.json.gz")
	topics, err := loadTopics(path)
	u.PanicIfErr(err)
	return topics
}

func loadTopicHistoriesMust() []TopicHistory {
	path := path.Join("stack-overflow-docs-dump", "topichistories.json.gz")
	topicHistories, err := loadTopicHistories(path)
	u.PanicIfErr(err)
	return topicHistories
}

func loadExamplesMust() []Example {
	path := path.Join("stack-overflow-docs-dump", "examples.json.gz")
	examples, err := loadExamples(path)
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
	gDocTags = loadDocTagsMust()
	gTopics = loadTopicsMust()
	gExamples = loadExamplesMust()
	gTopicHistories = loadTopicHistoriesMust()
	fmt.Printf("loadAll took %s\n", time.Since(timeStart))
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

func printIfNotEmpty(name, s string) {
	if len(s) == 0 {
		return
	}
	fmt.Printf("%s:\n%s\n\n", name, s)
}

func printTopic(t *Topic) {
	printIfNotEmpty("versions", t.VersionsJson)
	printIfNotEmpty("introduction", t.IntroductionMarkdown)
	printIfNotEmpty("parameters", t.ParametersMarkdown)
	printIfNotEmpty("remarks", t.RemarksMarkdown)
	printIfNotEmpty("syntax", t.SyntaxMarkdown)
}

/*
DocTopicHistoryTypeId int
DocTagId              int
DocTopicId            int
DocExampleId          int
*/

// TODO: speed up
func getExampleByID(id int) *Example {
	for i, e := range gExamples {
		if e.Id == id {
			return &gExamples[i]
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

func serFitsOneLine(s string) bool {
	if len(s) > 80 {
		return false
	}
	if strings.Contains(s, "\n") {
		return false
	}
	return true
}

func serField(k, v string) string {
	v = strings.TrimSpace(v)
	if len(v) == 0 {
		return ""
	}
	if serFitsOneLine(v) {
		return fmt.Sprintf("%s: %s\n", k, v)
	}
	return fmt.Sprintf("%s:\n%s\n---\n", k, v)
}

func shortenVersion(s string) string {
	if s == "[]" {
		return ""
	}
	return s
}

func writeIndexTxtMust(path string, topic *Topic) {
	s := serField("Title", topic.Title)
	s += serField("Versions", shortenVersion(topic.VersionsJson))
	s += serField("HtmlVersions", topic.HelloWorldVersionsHtml)
	s += serField("Introduction", topic.IntroductionMarkdown)
	s += serField("Syntax", topic.SyntaxMarkdown)
	s += serField("Parameters", topic.ParametersMarkdown)
	s += serField("Remarks", topic.RemarksMarkdown)

	err := ioutil.WriteFile(path, []byte(s), 0644)
	u.PanicIfErr(err)
	fmt.Printf("Wrote %s, %d bytes\n", path, len(s))
}

func writeSectionMust(path string, example *Example) {
	s := serField("Title", example.Title)
	s += "---\n"
	s += example.BodyMarkdown

	err := ioutil.WriteFile(path, []byte(s), 0644)
	u.PanicIfErr(err)
	fmt.Printf("Wrote %s, %d bytes\n", path, len(s))
}

func genBook(title string) {
	bookDir := makeURLSafe(title)
	docTag := findDocTagByTitleMust(gDocTags, title)
	//fmt.Printf("%s: docID: %d\n", title, docTag.Id)
	topics := getTopicsByDocTagID(docTag.Id)
	nChapters := len(topics)
	nSections := 0
	chapter := 10
	for _, t := range topics {
		examples := getExamplesForTopic(docTag.Id, t.Id)
		sortExamples(examples)

		dirChapter := fmt.Sprintf("%d-%s", chapter, makeURLSafe(t.Title))
		dirPath := filepath.Join("book", bookDir, dirChapter)
		err := os.MkdirAll(dirPath, 0755)
		u.PanicIfErr(err)
		chapterIndexPath := filepath.Join(dirPath, "index.txt")
		writeIndexTxtMust(chapterIndexPath, t)
		//fmt.Printf("%s\n", dirChapter)
		chapter += 10
		//fmt.Printf("%s, %d examples (%d), %s\n", t.Title, t.ExampleCount, len(examples), fileName)

		section := 10
		for _, ex := range examples {
			/*
				pinnedStr := ""
				if ex.IsPinned {
					pinnedStr = "pinned"
				}
			*/
			fileName := fmt.Sprintf("%d-%s.md", section, makeURLSafe(ex.Title))
			path := filepath.Join(dirPath, fileName)
			writeSectionMust(path, ex)
			//fmt.Printf("  %s %s '%s'\n", ex.Title, pinnedStr, fileName)
			//fmt.Printf("  %d-%s\n", section, fileName)
			//fmt.Printf("  %s\n", fileName)
			section += 10
		}
		nSections += len(examples)
	}
	fmt.Printf("%d chapters, %d sections\n", nChapters, nSections)
}

func main() {
	loadAll()
	genBook("jQuery")
}
