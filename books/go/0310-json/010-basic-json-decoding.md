---
Title: Basic JSON decoding
Id: 4113
Score: 4
---
[`json.Unmarshal`](https://golang.org/pkg/encoding/json/#Marshal) from the package `"encoding/json"` decodes a JSON value into the value pointed by the given variable.

The parameters are the value to decode in `[]bytes` and a variable to use as a storage for the de-serialized value. The returned value is an error (on failure).

```go
encodedValue := []byte(`{"London":18,"Rome":30}`)

// generic storage for the decoded JSON
var data map[string]interface{}

// decode the value into data
// notice that we must pass the pointer to data using &data
err := json.Unmarshal(encodedValue, &data)

// check if the decoding is successful
if err != nil {
    panic(err)
}

fmt.Println(data)
map[London:18 Rome:30]
```

[Playground](https://play.golang.org/p/CjplBCptH8)

Notice how in the example above we knew in advance both the type of the key and the value. But this is not always the case. In fact, in most cases the JSON contains mixed value types.

```go
encodedValue := []byte(`{"city":"Rome","temperature":30}`)

// generic storage for the decoded JSON
var data map[string]interface{}

// decode the value into data
if err := json.Unmarshal(encodedValue, &data); err != nil {
    panic(err)
}

// if you want to use a specific value type, we need to cast it
temp := data["temperature"].(float64)
fmt.Println(temp) // 30
city := data["city"].(string)
fmt.Println(city) // "Rome"
```

[Playground](https://play.golang.org/p/SawE86QKRt)

In the last example above we used a generic map to store the decoded value. We must use a `map[string]interface{}` because we know that the keys are strings, but we don't know the type of their values in advance.

This is a very simple approach, but it's also extremely limited. In the real world, you would generally [decode a JSON into a custom-defined `struct` type](a-22028).
