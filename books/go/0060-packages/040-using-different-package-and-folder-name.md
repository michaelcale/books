Title: Using different package and folder name
Id: 14559
Score: 0
Body:
It is perfectly fine to use a package name other than the folder name. If we do so, we still have to import the package based on the directory structure, but after the import we have to refer to it by the name we used in the package clause.

For example, if you have a folder `$GOPATH/src/mypck`, and in it we have a file `a.go`:

    package apple

    const Pi = 3.14

Using this package:

    package main

    import (
        "mypck"
        "fmt"
    )

    func main() {
        fmt.Println(apple.Pi)
    }

Even though this works, you should have a good reason to deviate package name from the folder name (or it may become source of misunderstanding and confusion).

### What's the use of this?

Simple. A package name is a Go [idetifier][2]:

    identifier = letter { letter | unicode_digit } .

Which allows unicode letters to be used in identifiers, e.g. `αβ` is a valid identifier in Go. Folder and file names are not handled by Go but by the Operating System, and different file systems have different restrictions. There are actually many file systems which would not allow all valid Go identifiers as folder names, so you would not be able to name your packages what otherwise the language spec would allow.

Having the option to use different package names than their containing folders, you have the option to really name your packages what the language spec allows, regardless of the underlying operating and file system.

  [1]: https://golang.org/ref/spec#Package_clause
  [2]: https://golang.org/ref/spec#Identifiers
|======|
