package main

import (
	"fmt"
	"log"
	"path"
	"sort"
	"time"

	"github.com/kjk/u"
)

var (
	gDocTags  []DocTag
	gTopics   []Topic
	gExamples []Example
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
	fmt.Printf("loadAll took %s\n", time.Since(timeStart))
}

func getTopicsByDocTagID(topics []Topic, docTagID int) []*Topic {
	var res []*Topic
	for i, topic := range topics {
		if topic.DocTagId == docTagID {
			res = append(res, &topics[i])
		}
	}
	return res
}

func genBook(title string) {
	docTag := findDocTagByTitleMust(gDocTags, title)
	fmt.Printf("%s: docID: %d\n", title, docTag.Id)

	topics := getTopicsByDocTagID(gTopics, docTag.Id)
	for i, t := range topics {
		fmt.Printf("%s, %d examples\n", t.Title, t.ExampleCount)
		if i == 0 {
			fmt.Printf("%s\n\n%s\n\n%s\n\n%s\n\n", t.IntroductionMarkdown, t.ParametersMarkdown, t.RemarksMarkdown, t.SyntaxMarkdown)
		}
	}
	fmt.Printf("%d topics\n", len(topics))
}

func main() {
	loadAll()
	genBook("jQuery")
}
