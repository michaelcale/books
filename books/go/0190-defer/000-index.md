---
Title: Defer
Id: 2795
---

In a complicated function, it's easy to miss releasing a resource like open file handle or to unlock a mutex.

In C++ you would use RAII technique. In Go you use `defer` statement.

```go
func foo() {
  f, err := os.Open("myfile.txt")
  if err != nil {
    return
  }
  defer f.Close()

  // ... lots of code
}
```

In the above example, `defer f.Close()` ensures that `f.Close()` will be called before we exit `foo`, even in the presence of a [panic](a-4350).

Placing `f.Close()` right after `os.Open()` makes it easier to audit the code and ensure `Close` is always called, even if there are multiple exit points in the function.

If deferred code is more complicated, you can use a function literal:

```go
func foo() {
  mutex1.Lock()
  mutex2.Lock()

  defer func() {
    mutex2.Unlock()
    mutex1.Unlock()
  }()

  // ... more code
}
```

You can have multiple `defer` statements. They'll be called in reverse order.
