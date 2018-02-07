---
Title: Maps
Search: dictionary, dictionaries, hash table, hash tables
Id: rd600050
---
A map is a mapping from value of one type to value of another type.

Another languages call them dictionaries (Python) or hash tables (C++).

Map basics:
```go
m := make(map[string]int)
m["number3"] = 3
k := "number3"
if n, ok := m[k]; ok {
    fmt.Printf("value of %s is %d\n", k, n)
} else {
    fmt.Printf("key '%s' doesn't exist in map m\n", k)
}
k = "number4"
if n, ok := m[k]; ok {
    fmt.Printf("value of %s is %d\n", k, n)
} else {
    fmt.Printf("key '%s' doesn't exist in map m\n", k)
}
```
**Output**:
```text
value of number3 is 3
key 'number4' doesn't exist in map m
```

[Zero value](a-6069) of map is nil.

Learn more about [maps](ch-732).


