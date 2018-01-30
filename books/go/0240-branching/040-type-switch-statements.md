Title: Type Switch Statements
Id: 4625
Score: 2
Body:
A simple type switch:
```
// assuming x is an expression of type interface{}
switch t := x.(type) {
case nil:
    // x is nil
    // t will be type interface{}
case int: 
    // underlying type of x is int
    // t will be int in this case as well
case string:
    // underlying type of x is string
    // t will be string in this case as well
case float, bool:
    // underlying type of x is either float or bool
    // since we don't know which, t is of type interface{} in this case
default:
    // underlying type of x was not any of the types tested for
    // t is interface{} in this type
}
```

----------

You can test for any type, including `error`, user-defined types, interface types, and function types:
```
switch t := x.(type) {
case error:
    log.Fatal(t)
case myType:
    fmt.Println(myType.message)
case myInterface:
    t.MyInterfaceMethod()
case func(string) bool:
    if t("Hello world?") {
        fmt.Println("Hello world!")
    }
}
```
|======|
