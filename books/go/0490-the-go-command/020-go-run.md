Title: Go Run
Id: 17005
Score: 2
Body:
`go run` will run a program without creating an executable file. Mostly useful for development. `run` will only execute packages whose *package name* is **main**.

To demonstrate, we will use a simple Hello World example `main.go`:

    package main
    
    import fmt
    
    func main() {
        fmt.Println("Hello, World!")
    }

Execute without compiling to a file:

    go run main.go

Output:

    Hello, World!

## Run multiple files in package

If the package is **main** and split into multiple files, one must include the other files in the `run` command:

    go run main.go assets.go

|======|
