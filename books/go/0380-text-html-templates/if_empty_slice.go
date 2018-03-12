package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
type UserTweets struct {
	User   string
	Tweets []string
}

const tmplStr = `
{{- if not .Tweets -}}
User '{{.User}}' has no tweets.

{{ else -}}
User '{{.User}}' has {{ len .Tweets }} tweets:
{{ range .Tweets -}}
  '{{ . }}'
{{ end }}
{{- end}}`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("if").Parse(tmplStr))

	data := UserTweets{
		User: "kjk",
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	data = UserTweets{
		User:   "masa",
		Tweets: []string{"tweet one", "tweet two"},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	// :show end
}
