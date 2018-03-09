---
Title: while loop
Id: 104
SOId: rd600078
---
Go doesn't have `while` loops.

`for` loop is powerful enough to express `while` loops.

For example, this C++ while loop:

```c++
int n = 0;
while (n < 3) {
    printf("n: %d\n", n);
    n++;
}
```

can be expressed as `for` loop:

```go
for n := 0; n < 3; n++ {
    fmt.Printf("n: %d\n", n)
}
```
