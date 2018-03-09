---
Title: Simple interface
Id: 91
SOId: 4001
---
In Go, an interface is just a set of methods. We use an interface to specify a behavior of a given object.

```go
type Painter interface {
    Paint()
}
```

The implementing type **need not** declare that it is implementing the interface. It is enough to define methods of the same signature.

```go
type Rembrandt struct{}

func (r Rembrandt) Paint() {
    // use a lot of canvas here
}
```

Now we can use the structure as an interface.

```
var p Painter
p = Rembrandt{}
```

An interface can be satisfied (or implemented) by an arbitrary number of types. Also a type can implement an arbitrary number of interfaces.

```go
type Singer interface {
        Sing()
}

type Writer interface {
        Write()
}

type Human struct{}

func (h *Human) Sing() {
    fmt.Println("singing")
}

func (h *Human) Write() {
    fmt.Println("writing")
}

type OnlySinger struct{}
func (o *OnlySinger) Sing() {
    fmt.Println("singing")
}
```

Here, The `Human` struct satisfy both the `Singer` and `Writer` interface, but the `OnlySinger` struct only satisfy `Singer` interface.
__________

**Empty Interface**

There is an empty interface type, that contains no methods. We declare it as `interface{}`. This contains no methods so every `type` satisfies it. Hence empty interface can contain any type value.

```go
var a interface{}
var i int = 5
s := "Hello world"

type StructType struct {
    i, j int
    k string
}

// all are valid statements
a = i
a = s
a = &StructType{1, 2, "hello"}
```

The most common use case for interfaces is to ensure that a variable supports one or more behaviours. By contrast, the primary use case for the empty interface is to define a variable which can hold any value, regardless of its concrete type.

To get these values back as their original types we just need to do

```go
i = a.(int)
s = a.(string)
m := a.(*StructType)
```

or

```go
i, ok := a.(int)
s, ok := a.(string)
m, ok := a.(*StructType)
```

`ok` indicates if the `interface a` is convertible to given type. If it is not possible to cast `ok` will be `false`.

___________

**Interface Values**

If you declare a variable of an interface, it may store any value type that implements the methods declared by the interface!

If we declare `h` of `interface Singer`, it may store a value of type `Human` or `OnlySinger.` This is because of the fact that they all implement methods specified by the `Singer` interface.

```go
var h Singer
h = &human{}

h.Sing()
```
