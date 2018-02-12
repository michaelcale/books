---
Title: Duplicate a struct (make a copy)
Id: 5352
Score: 4
---
A struct can simply be copied using assignment.

```go
type T struct {
    I int
    S string
}

// initialize a struct
t := T{1, "one"}

// make struct copy
u := t // u has its field values equal to t

if u == t { // true
    fmt.Println("u and t are equal") // Prints: "u and t are equal"
}
```

In above case, `'t'` and 'u' are now separate objects (struct values).

Since `T` does not contain any reference types (slices, map, channels) as its fields, `t` and `u` above can be modified without affecting each other.

```go
    fmt.Printf("t.I = %d, u.I = %d\n", t.I, u.I) // t.I = 100, u.I = 1
```

However, if `T` contains a reference type, for example:

```go
type T struct {
    I  int
    S  string
    xs []int // a slice is a reference type
}
```

Then a simple copy by assignment would copy the value of slice type field as well to the new object. This would result in two different objects referring to the same slice object.

```go
// initialize a struct
t := T{I: 1, S: "one", xs: []int{1, 2, 3}}

// make struct copy
u := t // u has its field values equal to t
```

Since both u and t refer to the same slice through their field xs updating a value in the slice of one object would reflect the change in the other.

```go
// update a slice field in u
u.xs[1] = 500

fmt.Printf("t.xs = %d, u.xs = %d\n", t.xs, u.xs)
// t.xs = [1 500 3], u.xs = [1 500 3]
```

Hence, extra care must be taken to ensure this reference type property does not produce unintended behavior.

To copy above objects for example, an explicit copy of the slice field could be performed:

```go
// explicitly initialize u's slice field
u.xs = make([]int, len(t.xs))
// copy the slice values over from t
copy(u.xs, t.xs)

// updating slice value in u will not affect t
u.xs[1] = 500

fmt.Printf("t.xs = %d, u.xs = %d\n", t.xs, u.xs)
// t.xs = [1 2 3], u.xs = [1 500 3]
```
