---
Title: Basic fmt
Id: 238
Score: 0
SOId: 23178
---
Package fmt implements formatted I/O using format *verbs*:

```text
%v    // the value in a default format
%T    // a Go-syntax representation of the type of the value
%s    // the uninterpreted bytes of the string or slice
```

# Format Functions

There are **4** main function types in `fmt` and several variations within.

## Print

```go
fmt.Print("Hello World")        // prints: Hello World
fmt.Println("Hello World")      // prints: Hello World\n
fmt.Printf("Hello %s", "World") // prints: Hello World
```

## Sprint

```go
formattedString := fmt.Sprintf("%v %s", 2, "words") // returns string "2 words"
```

## Fprint

```go
byteCount, err := fmt.Fprint(w, "Hello World") // writes to io.Writer w
```

`Fprint` can be used, inside `http` handlers:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s!", "Browser")
}   // Writes: "Hello Browser!" onto http response
```

## Scan

Scan scans text read from standard input.

```go
var s string
fmt.Scanln(&s) // pass pointer to buffer
// Scanln is similar to fmt.Scan(), but it stops scanning at new line.
fmt.Println(s) // whatever was inputted
```

# Stringer Interface

Any value which has a `String()` method implements the `fmt` **inteface** `Stringer`

```go
type Stringer interface {
        String() string
}
```
