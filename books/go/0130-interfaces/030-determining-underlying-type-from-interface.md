---
Title: Determining underlying type from interface
Id: 6081
Score: 2
---
Body:
In go it can sometimes be useful to know which underlying type you have been passed. This can be done with a type switch. This assumes we have two structs:
```go
type Rembrandt struct{}

func (r Rembrandt) Paint() {}

type Picasso struct{}

func (r Picasso) Paint() {}
```

That implement the Painter interface:
```go
type Painter interface {
    Paint()
}
```

Then we can use this switch to determine the underlying type:

```go
func WhichPainter(painter Painter) {
    switch painter.(type) {
    case Rembrandt:
        fmt.Println("The underlying type is Rembrandt")
    case Picasso:
        fmt.Println("The underlying type is Picasso")
    default:
        fmt.Println("Unknown type")
    }
}
```
