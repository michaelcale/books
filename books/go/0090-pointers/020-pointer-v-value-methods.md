Title: Pointer v. Value Methods
Id: 6049
Score: 1
Body:
# Pointer Methods

Pointer methods can be called even if the variable is itself not a pointer.

According to the [Go Spec](https://golang.org/ref/spec#Method_values),

>  . . . a reference to a non-interface method with a pointer receiver using an addressable value will automatically take the address of that value: `t.Mp` is equivalent to `(&t).Mp`.

You can see this in this example:

```
package main

import "fmt"

type Foo struct {
    Bar int
}

func (f *Foo) Increment() {
    f.Bar += 1
}

func main() {
    var f Foo

    // Calling `f.Increment` is automatically changed to `(&f).Increment` by the compiler.
    f = Foo{}
    fmt.Printf("f.Bar is %d\n", f.Bar)
    f.Increment()
    fmt.Printf("f.Bar is %d\n", f.Bar)
    
    // As you can see, calling `(&f).Increment` directly does the same thing.
    f = Foo{}
    fmt.Printf("f.Bar is %d\n", f.Bar)
    (&f).Increment()
    fmt.Printf("f.Bar is %d\n", f.Bar)
}
```
**[Play it](https://play.golang.org/p/jlQLrSnH-E)**

# Value Methods

Similarly to pointer methods, value methods can be called even if the variable is itself not a value.

According to the [Go Spec](https://golang.org/ref/spec#Method_values),

>  . . . a reference to a non-interface method with a value receiver using a pointer will automatically dereference that pointer: `pt.Mv` is equivalent to `(*pt).Mv`.

You can see this in this example:

```
package main

import "fmt"

type Foo struct {
    Bar int
}

func (f Foo) Increment() {
    f.Bar += 1
}

func main() {
    var p *Foo

    // Calling `p.Increment` is automatically changed to `(*p).Increment` by the compiler.
    // (Note that `*p` is going to remain at 0 because a copy of `*p`, and not the original `*p` are being edited)
    p = &Foo{}
    fmt.Printf("(*p).Bar is %d\n", (*p).Bar)
    p.Increment()
    fmt.Printf("(*p).Bar is %d\n", (*p).Bar)
    
    // As you can see, calling `(*p).Increment` directly does the same thing.
    p = &Foo{}
    fmt.Printf("(*p).Bar is %d\n", (*p).Bar)
    (*p).Increment()
    fmt.Printf("(*p).Bar is %d\n", (*p).Bar)
}
```
**[Play it](https://play.golang.org/p/Efc0IVgzh8)**

----------

To learn more about pointer and value methods, visit the [Go Spec section on Method Values](https://golang.org/ref/spec#Method_values), or see the [Effective Go section about Pointers v. Values](https://golang.org/doc/effective_go.html#pointers_vs_values).

----------

_Note 1: The parenthesis (`()`) around `*p` and `&f` before selectors like `.Bar` are there for grouping purposes, and must be kept._

_Note 2: Although pointers can be converted to values (and vice-versa) when they are the receivers for a method, they are_ not _automattically converted to eachother when they are arguments inside of a function._
|======|
