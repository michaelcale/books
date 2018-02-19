---
Title: Basic JSON encoding
Search: JSON serialize
Id: 4112
Score: 1
---
[`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) from the package `"encoding/json"` encodes a value to JSON.

The parameter is the value to encode. The returned values are an array of bytes representing the JSON-encoded input (on success), and an error (on failure).

```go
decodedValue := []string{"foo", "bar"}

// encode the value
data, err := json.Marshal(decodedValue)

// check if the encoding is successful
if err != nil {
    panic(err)
}

// print out the JSON-encoded string
// remember that data is a []byte
fmt.Println(string(data))
// "["foo","bar"]"
```

Here's some basic examples of encoding for built-in data types:

```go
var data []byte

data, _ = json.Marshal(1)
fmt.Println(string(data))
// 1

data, _ = json.Marshal("1")
fmt.Println(string(data))
// "1"

data, _ = json.Marshal(true)
fmt.Println(string(data))
// true

data, _ = json.Marshal(map[string]int{"London": 18, "Rome": 30})
fmt.Println(string(data))
// {"London":18,"Rome":30}
```

Encoding simple variables is helpful to understand how the JSON encoding works in Go. However, in the real world, you'll likely [encode more complex data stored in structs](a-22028).
