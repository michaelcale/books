---
Title: Handling an error
Id: 124
SOId: 2708
---
In Go errors can be returned from a function call. The convention is that if a method can fail, the last returned argument is an `error`.

```go
func DoAndReturnSomething() (string, error) {
    if os.Getenv("ERROR") == "1" {
        return "", errors.New("The method failed")
    }

    // The method succeeded.
    return "Success!", nil
}
```

You use multiple variable assignments to check if the method failed.

```go
result, err := DoAndReturnSomething()
if err != nil {
    panic(err)
}

// This is executed only if the method didn't return an error
fmt.Println(result)
```

If you are not interested in the error, you can simply ignore it by assigning it to `_`.

```go
result, _ := DoAndReturnSomething()
fmt.Println(result)
```

Of course, ignoring an error can have serious implications. Therefore, this is generally not recommended.

If you have multiple method calls, and one or more methods in the chain may return an error, you should propagate the error to the first level that can handle it.

```go
func Foo() error {
    return errors.New("I failed!")
}

func Bar() (string, error) {
    err := Foo()
    if err != nil {
        return "", err
    }

    return "I succeeded", nil
}

func Baz() (string, string, error) {
    res, err := Bar()
    if err != nil {
        return "", "", err
    }

    return "Foo", "Bar", nil
}
```
