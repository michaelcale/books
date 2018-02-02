Title: Using bytes.Reader to read from a string
Id: 23599
Score: 0
Body:
One implementation of the `io.Reader` interface can be found in the `bytes` package. It allows a byte slice to be used as the source for a Reader. In this example the byte slice is taken from a string, but is more likely to have been read from a file or network connection.
~~~
message := []byte("Hello, playground")

reader := bytes.NewReader(message)

bs := make([]byte, 5)
n, err := reader.Read(bs)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Read %d bytes: %s", n, bs)
~~~

[Go Playground][1]

  [1]: https://play.golang.org/p/cRSRKwKcXr
|======|
