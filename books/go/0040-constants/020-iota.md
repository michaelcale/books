---
Title: iota
Id: 80100036
---
`iota` makes it easy to declare number constants whose values grow.

Basics:
```go
const (
    Low = iota
    Medium
    High
)
fmt.Printf("Low: %d\nMedium: %d\nHigh: %d\n", Low, Medium, High)
```
**Output**:
```text
Low: 0
Medium: 1
High: 2
```

`iota` sets the value of `Low` to 0 and instructs the compiler following constants have increasing numeric values.


## Creating bitmask values with `iota`

Iota can be very useful when creating a bitmask. For instance, to represent the state of a network connection which may be secure, authenticated, and/or ready, we might create a bitmask like the following:

```go
const (
    Secure = 1 << iota // 0b001
    Authn              // 0b010
    Ready              // 0b100
)

ConnState := Secure|Authn // 0b011: Connection is secure and authenticated, but not yet Ready
```

## Skipping values

The value of `iota` is still incremented for every entry in a constant list even if iota is not used:

```go
const ( // iota is reset to 0
    a = 1 << iota  // a == 1
    b = 1 << iota  // b == 2
    c = 3          // c == 3  (iota is not used but still incremented)
    d = 1 << iota  // d == 8
)
```

it will also be incremented even if no constant is created at all, meaning the empty identifier can be used to skip values entirely:

```go
const (
    a = iota // a = 0
    _        // iota is incremented
    b        // b = 2
)
```

## Using `iota` in an expression list

Because `iota` is incremented after each [`ConstSpec`](https://golang.org/ref/spec#ConstSpec), values within the same expression list will have the same value for `iota`:

```go
const (
    bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
    bit1, mask1                           // bit1 == 2, mask1 == 1
    _, _                                  // skips iota == 2
    bit3, mask3                           // bit3 == 8, mask3 == 7
)
```

## Using `iota` in an expression

`iota` can be used in expressions, so it can also be used to assign values other than simple incrementing integers starting from zero. To create constants for SI units, use this example from [Effective Go](https://golang.org/doc/effective_go.html#initialization):

```go
type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```
