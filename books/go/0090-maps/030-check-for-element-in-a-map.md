---
Title: Get value by key
Id: 4650
---
To get a value from the map, you just have to do something like:00

    value := mapName[ key ]

If the map contains the key, it returns the corresponding value.
If not, it returns zero-value of the map's value type (`0` if map of `int` values, `""` if map of `string` values...)

```go
m  := map[string]string{"foo": "foo_value", "bar": ""}
k  := m["foo"]  // returns "foo_value" since that is the value stored in the map
k2 := m["bar"]  // returns "" since that is the value stored in the map
k3 := m["nop"]  // returns "" since the key does not exist, and "" is the string type's zero value
```

To differentiate between empty values and non-existent keys, you can use the second returned value of the map access (for example `value, hasKey := map["key"]`).

This second value is `boolean` typed, and will be:
- `true` when the value exists in the map,
- `false` when the map does not contains the given key.

Look at the following example:

```go
value, hasKey = m[ key ]
if hasKey {
    // the map contains the given key, so we can safely use the value
    // If value is zero-value, it's because the zero-value was pushed to the map
} else {
    // The map does not have the given key
    // the value will be the zero-value of the map's type
}
```
