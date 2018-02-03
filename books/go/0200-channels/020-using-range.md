Title: Using range
Id: 4134
Score: 2
Body:
When reading multiple values from a channel, using `range` is a common pattern:

```
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

[Playground][1]

Output

```
1
2
3
channel is now closed
```


  [1]: https://play.golang.org/p/18ODvaZub9
|======|
