package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
var tmplStr = `User {{.User}} has {{.TotalTweets}} tweets.
{{- $tweetCount := len .RecentTweets }}
Recent tweets:
{{range $idx, $tweet := .RecentTweets}}Tweet {{$idx}} of {{$tweetCount}}: '{{.}}'
{{end -}}
Most recent tweet: '{{index .RecentTweets 0}}'`

// :show end

func main() {
	// :show start
	t := template.New("tweets")
	t, err := t.Parse(tmplStr)
	if err != nil {
		log.Fatalf("template.Parse() failed with '%s'\n", err)
	}
	data := struct {
		User         string
		TotalTweets  int
		RecentTweets []string
	}{
		User:         "kjk",
		TotalTweets:  124,
		RecentTweets: []string{"hello", "there"},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
