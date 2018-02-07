---
Title: Parameters
Id: 1251
Score: 0
---
A function can optionally declare a set of parameters:

```go
func SayHelloToMe(firstName, lastName string, age int) {
    fmt.Printf("Hello, %s %s!\n", firstName, lastName)
    fmt.Printf("You are %d", age)
}
```

Notice that the type for `firstName` is omitted because it is identical to `lastName`.
