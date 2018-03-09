---
Title: Testing HTTP requests
Id: 249
Score: 1
SOId: 9053
---
`main.go`:

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func fetchContent(url string) (string, error) {
    res, err := http.Get(url)
    if err != nil {
        return "", nil
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}

func main() {
    url := "https://example.com/"
    content, err := fetchContent(url)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Content:", content)
}
```

`main_test.go`:

```go
package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Test_fetchContent(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "hello world")
    }))
    defer ts.Close()

    content, err := fetchContent(ts.URL)
    if err != nil {
        t.Error(err)
    }

    want := "hello world"
    if content != want {
        t.Errorf("Got %q, want %q", content, want)
    }
}
```
