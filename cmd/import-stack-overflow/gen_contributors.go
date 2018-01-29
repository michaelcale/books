package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kjk/u"
)

/*
ad-hoc code to create a list of stack overflow contributors for each book.
We only have user ids and we can convert them to names with:

$ http HEAD https://stackoverflow.com/users/1850609/
HTTP/1.1 301 Moved Permanently
Location: /users/1850609/acdcjunior
*/

func genContributorsAndExit() {
	timeStart := time.Now()
	fmt.Printf("Loading Stack Overflow data...")
	gDocTags = loadDocTagsMust()
	gContributors = loadContributorsMust()
	gTopics = loadTopicsMust()
	gExamples = loadExamplesMust()
	fmt.Printf(" took %s\n", time.Since(timeStart))

	// for fast search, build topic id => doc id mapping
	topicToBook := make(map[int]int)
	for _, topic := range gTopics {
		topicToBook[topic.Id] = topic.DocTagId
	}

	nMissingTopics := 0
	nMissingExamples := 0

	// for fast search, build example id => doc id mapping
	exampleToBook := make(map[int]int)
	for _, example := range gExamples {
		exampleID := example.Id
		topicID := example.DocTopicId
		bookID, ok := topicToBook[topicID]
		if !ok {
			//u.PanicIf(!ok, "missing topic id => book id in topicToBook, topicID: %d", topicID)
			nMissingTopics++
		}
		exampleToBook[exampleID] = bookID
	}

	// TODO: build doc id = > doc name mapping

	// maps book id to map of contributor id => count of contributions
	perBookContributors := make(map[int]map[int]int)

	contributors := make(map[int]int)
	var bookID int
	var ok bool
	for _, c := range gContributors {
		contributors[c.UserId]++
		topicID := c.DocTopicId
		if topicID != 0 {
			bookID, ok = topicToBook[topicID]
			if !ok {
				//u.PanicIf(!ok, "missing topic id = > book id in topicToBook, topicID: %d", topicID)
				nMissingTopics++
			}
		} else {
			exampleID := c.DocExampleId
			u.PanicIf(exampleID == 0, "both topicID and exampleID is 0")
			bookID, ok = exampleToBook[exampleID]
			if !ok {
				//u.PanicIf(!ok, "missing example id = > book id in exampleToBook, exampleID: %d", exampleID)
				nMissingExamples++
			}
		}
		bookContrib, ok := perBookContributors[bookID]
		if !ok {
			bookContrib = make(map[int]int)
			perBookContributors[bookID] = bookContrib
		}
		bookContrib[c.UserId]++
	}

	fmt.Printf("Total contributions: %d\n", len(gContributors))
	fmt.Printf("Unique contributors: %d\n", len(contributors))
	fmt.Printf("Total books: %d\n", len(perBookContributors))
	fmt.Printf("Missing topics: %d\n", nMissingTopics)
	fmt.Printf("Missing examples: %d\n", nMissingExamples)
	max := 32
	for cID, n := range contributors {
		fmt.Printf("%d, %d\n", cID, n)
		max--
		if max < 0 {
			break
		}
	}
	os.Exit(0)
}
