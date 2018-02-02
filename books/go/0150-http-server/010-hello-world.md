Title: Hello World
Id: 2703
Score: 13
Body:
The typical way to begin writing webservers in golang is to use the standard library `net/http` module.  

There is also a tutorial for it [here][1].  

The following code also uses it.  Here is the simplest possible HTTP server implementation. It responds `"Hello World"` to any HTTP request.

Save the following code in a `server.go` file in your workspaces.

    package main
    
    import (
        "log"
        "net/http"
    )
    
    func main() {
        // All URLs will be handled by this function
        // http.HandleFunc uses the DefaultServeMux
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Hello, world!"))
        })
    
        // Continue to process new requests until an error occurs
        log.Fatal(http.ListenAndServe(":8080", nil))
    }

You can run the server using:

    $ go run server.go

Or you can compile and run.

    $ go build server.go
    $ ./server

The server will listen to the specified port (`:8080`). You can test it with any HTTP client. Here's an example with `cURL`:

    curl -i http://localhost:8080/
    HTTP/1.1 200 OK
    Date: Wed, 20 Jul 2016 18:04:46 GMT
    Content-Length: 13
    Content-Type: text/plain; charset=utf-8
    
    Hello, world!

Press <kbd>Ctrl</kbd>+<kbd>C</kbd> to stop the process.


  [1]: https://golang.org/doc/articles/wiki/
|======|
