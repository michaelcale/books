Title: Go Fmt
Id: 17008
Score: 3
Body:
`go fmt` will format a program's source code in a neat, idiomatic way that is easy to read and understand. It is recommended that you use `go fmt` on any source before you submit it for public viewing or committing into a version control system, to make reading it easier.

To format a file:

    go fmt main.go

Or all files in a directory:

    go fmt myProject

You can also use `gofmt -s` (**not** `go fmt`) to attempt to simplify any code that it can.

`gofmt` (**not** `go fmt`) can also be used to refactor code. It understands Go, so it is more powerful than using a simple search and replace. For example, given this program (`main.go`):

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

You can replace the method `Original` with `Refactor` with `gofmt`:

    gofmt -r 'Original -> Refactor' -d main.go

Which will produce the diff:

    -func (e *Example) Original(name string) {
    +func (e *Example) Refactor(name string) {
         e.Name = name
     }
     
     func main() {
         e := &Example{"Hello"}
    -    e.Original("Goodbye")
    +    e.Refactor("Goodbye")
     }
|======|
