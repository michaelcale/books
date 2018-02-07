---
Title: Using a handler function
Id: 2704
Score: 0
---
[`HandleFunc`](https://golang.org/pkg/net/http/#ListenAndServe) registers the handler function for the given pattern in the server mux (router).

You can pass define an anonymous function, as we have seen in the basic _Hello World_ example:

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}
```

But we can also pass a [`HandlerFunc`](https://golang.org/pkg/net/http/#HandlerFunc) type. In other words, we can pass any function that respects the following signature:

```go
func FunctionName(w http.ResponseWriter, req *http.Request)
```

We can rewrite the previous example passing the reference to a previously defined `HandlerFunc`. Here's the full example:

```go
package main

import (
    "fmt"
    "net/http"
)

// A HandlerFunc function
// Notice the signature of the function
func RootHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}

func main() {
    // Here we pass the reference to the `RootHandler` handler function
    http.HandleFunc("/", RootHandler)
    panic(http.ListenAndServe(":8080", nil))
}
```

Of course, you can define several function handlers for different paths.

```
package main

import (
    "fmt"
    "log"
    "net/http"
)

func FooHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Hello from foo!")
}

func BarHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Hello from bar!")
}

func main() {
    http.HandleFunc("/foo", FooHandler)
    http.HandleFunc("/bar", BarHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Here's the output using `cURL`:

```bash
➜  ~ curl -i localhost:8080/foo
HTTP/1.1 200 OK
Date: Wed, 20 Jul 2016 18:23:08 GMT
Content-Length: 16
Content-Type: text/plain; charset=utf-8

Hello from foo!

➜  ~ curl -i localhost:8080/bar
HTTP/1.1 200 OK
Date: Wed, 20 Jul 2016 18:23:10 GMT
Content-Length: 16
Content-Type: text/plain; charset=utf-8

Hello from bar!

➜  ~ curl -i localhost:8080/
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 20 Jul 2016 18:23:13 GMT
Content-Length: 19

404 page not found
```
