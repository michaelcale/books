---
Title: Different Forms of For Loop
Id: 7956
Score: 3
---
**Simple form using one variable:**

```go
for i := 0; i < 10; i++ {
    fmt.Print(i, " ")
}
```

**Using two variables (or more):**

```go
for i, j := 0, 0; i < 5 && j < 10; i, j = i+1, j+2 {
    fmt.Println(i, j)
}
```

**Without using initialization statement:**

```go
i := 0
for ; i < 10; i++ {
    fmt.Print(i, " ")
}
```

**Without a test expression:**

```go
for i := 1; ; i++ {
    if i&1 == 1 {
        continue
    }
    if i == 22 {
        break
    }
    fmt.Print(i, " ")
}
```

**Without increment expression:**

```go
for i := 0; i < 10; {
    fmt.Print(i, " ")
    i++
}
```

**When all three initialization, test and increment expressions are removed, the loop becomes infinite:**

```go
i := 0
for {
    fmt.Print(i, " ")
    i++
    if i == 10 {
        break
    }
}
```

**This is an example of infinite loop with counter initialized with zero:**

```go
for i := 0; ; {
    fmt.Print(i, " ")
    if i == 9 {
        break
    }
    i++
}
```

**When just the test expression is used (acts like a typical while loop):**

```go
i := 0
for i < 10 {
    fmt.Print(i, " ")
    i++
}
```

**Using just increment expression:**

```go
i := 0
for ; ; i++ {
    fmt.Print(i, " ")
    if i == 9 {
        break
    }
}
```

**Iterate over a range of values using index and value:**

```go
ary := [5]int{0, 1, 2, 3, 4}
for index, value := range ary {
    fmt.Println("ary[", index, "] =", value)
}
```

**Iterate over a range using just index:**

```go
for index := range ary {
    fmt.Println("ary[", index, "] =", ary[index])
}
```

**Iterate over a range using just index:**

```go
for index, _ := range ary {
    fmt.Println("ary[", index, "] =", ary[index])
}
```

**Iterate over a range using just value:**

```go
for _, value := range ary {
    fmt.Print(value, " ")
}
```

**Iterate over a range using key and value for map (may not be in order):**

```go
mp := map[string]int{"One": 1, "Two": 2, "Three": 3}
for key, value := range mp {
    fmt.Println("map[", key, "] =", value)
}
```

**Iterate over a range using just key for map (may be not in order):**

```go
for key := range mp {
    fmt.Print(key, " ") //One Two Three
}
```

**Iterate over a range using just key for map (may be not in order):**

```go
for key, _ := range mp {
    fmt.Print(key, " ") //One Two Three
}
```

**Iterate over a range using just value for map (may be not in order):**

```go
for _, value := range mp {
    fmt.Print(value, " ") //2 3 1
}
```

**Iterate over a range for channels (exits if the channel is closed):**

```go
ch := make(chan int, 10)
for i := 0; i < 10; i++ {
    ch <- i
}
close(ch)

for i := range ch {
    fmt.Print(i, " ")
}
```

**Iterate over a range for string (gives Unicode code points):**

```go
utf8str := "B = \u00b5H" //B = µH
for _, r := range utf8str {
    fmt.Print(r, " ") //66 32 61 32 181 72
}
fmt.Println()
for _, v := range []byte(utf8str) {
    fmt.Print(v, " ") //66 32 61 32 194 181 72
}
fmt.Println(len(utf8str)) //7
```

as you see `utf8str` has 6 runes (Unicode code points) and 7 bytes.
