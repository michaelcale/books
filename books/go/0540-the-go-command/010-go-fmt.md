Title: go fmt
Search: gofmt
Id: 17008
Body:
To keep code consistent and eliminate arguments over code formatting, Go includes `go fmt` tool.

To format a file: `go fmt main.go`

Or all files in a directory: `go fmt myProject`

You can also use `gofmt -s` (**not** `go fmt`) to attempt to simplify code whenever possible.

`gofmt` (**not** `go fmt`) can also be used to refactor code. It understands Go syntax so it is more powerful than a search and replace. For example, given this program (`main.go`):

```go
package main

type Example struct {
    Name string
}

func (e *Example) Original(name string) {
    e.Name = name
}

func main() {
    e := &Example{"Hello"}
    e.Original("Goodbye")
}
```

You can replace the method `Original` with `Refactor` with `gofmt`:

```bash
$ gofmt -r 'Original -> Refactor' -d main.go
```

Which will produce the diff:

```diff
-func (e *Example) Original(name string) {
+func (e *Example) Refactor(name string) {
        e.Name = name
    }

    func main() {
        e := &Example{"Hello"}
-    e.Original("Goodbye")
+    e.Refactor("Goodbye")
    }
```
