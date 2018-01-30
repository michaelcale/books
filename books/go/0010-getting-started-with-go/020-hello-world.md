Title: Hello, World!
Id: 833
Score: 92
Body:
Place the following code into a file name `hello.go`:

    package main
    
    import "fmt"
    
    func main() {
        fmt.Println("Hello, 世界")
    }

[Playground](https://play.golang.org/p/I3l_5RKJts)

When `Go` is [installed correctly][1] this program can be compiled and run like this:

    go run hello.go

# Output:

```
Hello, 世界
```

Once you are happy with the code it can be compiled to an executable by running:

    go build hello.go

This will create  an executable file appropriate for your operating system in the current directory, which you can then run with the following command:

**Linux, OSX, and other Unix-like systems**

    ./hello

**Windows**

    hello.exe


_**Note**: The Chinese characters are important because they demonstrate that Go strings are stored as read-only slices of bytes._


  [1]: https://golang.org/dl/
|======|
