Title: Basic fmt
Id: 23178
Score: 0
Body:
Package fmt implements formatted I/O using format *verbs*:

    %v    // the value in a default format
    %T    // a Go-syntax representation of the type of the value
    %s    // the uninterpreted bytes of the string or slice

# Format Functions 

There are **4** main function types in `fmt` and several variations within.

## Print

    fmt.Print("Hello World")        // prints: Hello World
    fmt.Println("Hello World")      // prints: Hello World\n
    fmt.Printf("Hello %s", "World") // prints: Hello World

## Sprint

    formattedString := fmt.Sprintf("%v %s", 2, "words") // returns string "2 words"

## Fprint

    byteCount, err := fmt.Fprint(w, "Hello World") // writes to io.Writer w
   
`Fprint` can be used, inside `http` handlers:

    func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello %s!", "Browser")
    }   // Writes: "Hello Browser!" onto http response

## Scan

Scan scans text read from standard input.

    var s string
    fmt.Scanln(&s) // pass pointer to buffer
    // Scanln is similar to fmt.Scan(), but it stops scanning at new line.
    fmt.Println(s) // whatever was inputted

# Stringer Interface

Any value which has a `String()` method implements the `fmt` **inteface** `Stringer`

    type Stringer interface {
            String() string
    }
|======|
