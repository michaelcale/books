---
Title: Split and join strings
Id: 5eq100s7
---

## Split a string using `strings.Split`

TODO: write me

## Join a string using `strings.Join`

TODO: write me

* [`strings.Split`](https://golang.org/pkg/strings/#Split)

      s := "foo, bar, bar"
      fmt.Println(strings.Split(s, ", ")) // [foo bar baz]

* [`strings.Join`](https://golang.org/pkg/strings/#Join)

      ss := []string{"foo", "bar", "bar"}
      fmt.Println(strings.Join(ss, ", ")) // foo, bar, baz

