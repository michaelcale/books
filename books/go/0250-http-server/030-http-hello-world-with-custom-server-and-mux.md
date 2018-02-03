Title: HTTP Hello World with custom server and mux
Id: 2570
Score: 4
Body:
    package main
    
    import (
        "log"
        "net/http"
    )
    
    func main() {

        // Create a mux for routing incoming requests
        m := http.NewServeMux()

        // All URLs will be handled by this function
        m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Hello, world!"))
        })

        // Create a server listening on port 8000
        s := &http.Server{
            Addr:    ":8000",
            Handler: m,
        }

        // Continue to process new requests until an error occurs
        log.Fatal(s.ListenAndServe())
    }

Press <kbd>Ctrl</kbd>+<kbd>C</kbd> to stop the process.
|======|
