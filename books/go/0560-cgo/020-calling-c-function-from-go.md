Title: Calling C Function From Go
Id: 22191
Score: 3
Body:
Cgo enables the creation of Go packages that call C code.  
To use `cgo` write normal Go code that imports a pseudo-package "C". The Go code can then refer to types such as `C.int`, or functions such as `C.Add`.  
The import of "C" is immediately preceded by a comment, that comment, called the preamble, is used as a header when compiling the C parts of the package.  
Note that there must be no blank lines in between the `cgo` comment and the import statement.  
Note that `import "C"` can not  grouped with other imports into a parenthesized, "factored" import statement. You must write multiple import statements, like:

    import "C"
    import "fmt"

And it is good style to use the factored import statement, for other imports, like:

    import "C"
    import (
        "fmt"
        "math"
    )

Simple example using `cgo`:

    package main
    
    //int Add(int a, int b){
    //    return a+b;
    //}
    import "C"
    import "fmt"
    
    func main() {
        a := C.int(10)
        b := C.int(20)
        c := C.Add(a, b)
        fmt.Println(c) // 30
    }
Then `go build`, and run it, output:  

    30

To build `cgo` packages, just use `go build` or `go install` as usual. The `go tool` recognizes the special `"C"` import and automatically uses `cgo` for those files.

|======|
