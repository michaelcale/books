---
Title: strings package
Id: 30331
Score: 1
---
* [`strings.Contains`](https://golang.org/pkg/strings/#Contains)

      fmt.Println(strings.Contains("foobar", "foo")) // true
      fmt.Println(strings.Contains("foobar", "baz")) // false

* [`strings.HasPrefix`](https://golang.org/pkg/strings/#HasPrefix)

      fmt.Println(strings.HasPrefix("foobar", "foo")) // true
      fmt.Println(strings.HasPrefix("foobar", "baz")) // false

* [`strings.HasSuffix`](https://golang.org/pkg/strings/#HasSuffix)

      fmt.Println(strings.HasSuffix("foobar", "bar")) // true
      fmt.Println(strings.HasSuffix("foobar", "baz")) // false

* [`strings.Join`](https://golang.org/pkg/strings/#Join)

      ss := []string{"foo", "bar", "bar"}
      fmt.Println(strings.Join(ss, ", ")) // foo, bar, baz

* [`strings.Replace`](https://golang.org/pkg/strings/#Replace)

      fmt.Println(strings.Replace("foobar", "bar", "baz", 1)) // foobaz

* [`strings.Split`](https://golang.org/pkg/strings/#Split)

      s := "foo, bar, bar"
      fmt.Println(strings.Split(s, ", ")) // [foo bar baz]

* [`strings.ToLower`](https://golang.org/pkg/strings/#ToLower)

      fmt.Println(strings.ToLower("FOOBAR")) // foobar

* [`strings.ToUpper`](https://golang.org/pkg/strings/#ToUpper)

      fmt.Println(strings.ToUpper("foobar")) // FOOBAR

* [`strings.TrimSpace`](https://golang.org/pkg/strings/#TrimSpace)

      fmt.Println(strings.TrimSpace("  foobar  ")) // foobar

More: https://golang.org/pkg/strings/.
