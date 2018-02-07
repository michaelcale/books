---
Title: Literal functions & closures
Id: 1265
Score: 7
---
A simple literal function, printing `Hello!` to stdout:

```go
package main

import "fmt"

func main() {
    func(){
        fmt.Println("Hello!")
    }()
}
```

[play it on playground](https://play.golang.org/p/upOAwpOaue)

----------

A literal function, printing the `str` argument to stdout:

```go
package main

import "fmt"

func main() {
    func(str string) {
        fmt.Println(str)
    }("Hello!")
}
```

[play it on playground](https://play.golang.org/p/jz-5wpEkY2)

----------

A literal function, closing over the variable `str`:

```go
package main

import "fmt"

func main() {
    str := "Hello!"
    func() {
        fmt.Println(str)
    }()
}
```

[play it on playground](https://play.golang.org/p/j6ZgyAna7l)

----------

It is possible to assign a literal function to a variable:

```go
package main

import (
    "fmt"
)

func main() {
    str := "Hello!"
    anon := func() {
        fmt.Println(str)
    }
    anon()
}
```

[play it on playground](https://play.golang.org/p/Ick7RmdTFb)
