---
Title: Using range
Id: 4134
Score: 2
---
When reading multiple values from a channel, using `range` is a common pattern:

```go
func foo() chan int {
    ch := make(chan int)

    go func() {
        ch <- 1
        ch <- 2
        ch <- 3
        close(ch)

    }()

    return ch
}

func main() {
    for n := range foo() {
        fmt.Println(n)
    }

    fmt.Println("channel is now closed")
}
```

[Playground](https://play.golang.org/p/18ODvaZub9)

Output

```text
1
2
3
channel is now closed
```
