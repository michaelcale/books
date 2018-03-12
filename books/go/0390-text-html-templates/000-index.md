---
Title: Text and HTML templates
Id: 217
SOId: 3888
---

Imagine you're working on a web application and need to return HTML that lists most recent tweets. You need to load list of tweets from a database and create HTML based on that information.

Building that HTML string by building smaller strings and concatenating them with `+` would be tedious.

Packages `text/template` and `html/template` in Go standard library make implement data-driven templates for generating textual output:

@file index.go output

Each template has a name. `template.New("tweets")` creates an empty template with name `tweets`.

`t.Parse(s string)` parses the template.

`t.Execute(w io.Writer, v interface{})` executes the template with a given value and writing the result to an `io.Writer`.

`{{ ... }}` is an action are instructions for templating engine.

`{{.TweetCount}}` means printing the value of `TweetCount` in current context.

Data passed to a template can be hierarchical (i.e. a struct withing a struct within a struct...).

Current context `.` refers to current scope within the data.

Initial `.` refers to top-level scope:

@file index2.go output

Values that don't have pre-defined formatting are printed using `Stringer` interface. For custom formatting of your type in a template implement `String() string` method.

`{{range .Tweets}}{{end}}` evaluates inner part for every element of `[]string` slice `Tweets` and sets current context `.` within the inner part to elements of `Tweets` slice.

`{{index .RecentTweets 0}}` is equivalent to `RecentTweets[0]` in Go code.

Text in a template is copied verbatim. Having to preserve whitespace can lead to ugly templates.

To help write more readable templates We can add `-` at the beginning or end of action as seen in `{{end -}}`.

This remove whitespace before or after the action.

`{{range .RecentTweets}}` changes variable scope and we don't have access to data outside. If we need to access data from upper scope, we can define variables like `{{ $tweetCount := len .RecentTweets }}`.
