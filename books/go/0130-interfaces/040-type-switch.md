---
Title: Type switch
Id: 14736
Score: 1
---
Type switches can also be used to get a variable that matches the type of the case:

```go
func convint(v interface{}) (int,error) {
    switch u := v.(type) {
    case int:
        return u, nil
    case float64:
        return int(u), nil
    case string:
        return strconv(u)
    default:
        return 0, errors.New("Unsupported type")
    }
}
```
