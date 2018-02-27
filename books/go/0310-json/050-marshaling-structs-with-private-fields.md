---
Title: Custom JSON marshaling
Id: 14194
---

## Writing custom JSON marshalling

Sometimes a type doesn't have an obvious mapping to JSON.

How would you serialize `time.Time`? There are so many possibilities.

Go provides a default JSON mapping for `time.Time`. We can implement custom marshaller for user-defined types like structs.

For existing types we can define a new (but compatible) type.

Here's a custom marshalling and unmarshaling for `time.Time` that only serializes year/month/date part:

@file custom_marshal.go output sha1:2d8f94c56472607c83551b5df2741bddb7451a6d goplayground:QpOf0zMOctW

Notice that receiver type of `UnmashalJSON` is a pointer to the type.

This is necessary for changes to persist outside the function itself.

## Marshaling structs with private fields

Consider a struct with both exported and unexported fields:

```go
type MyStruct struct {
    uuid string
    Name string
}
```

Imagine you want to `Marshal()` this struct into valid JSON for storage in something like etcd.

However, since `uuid` in not exported, the `json.Marshal()` skips it.

To marshal private fields without making them public we can use a custom marshaller:

@file marshal_private.go output sha1:aaae4444a438ff81e36995959a675309c03890b2 goplayground:JLd2aY-_5sh

## Custom marshaling behind the scenes

How does custom marshaling works?

Package JSON defines 2 interfaces: [`Marshaler`](https://golang.org/pkg/encoding/json/#Marshaler) and [`Unmarshaler`](https://golang.org/pkg/encoding/json/#Unmarshaler).

```go
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}
```

By implementing those functions we make our type conform to `Marshaler` or `Unmarshaler` interface.

JSON encoder / decoder checks if the value being encoded conforms to those interfaces and will call those functions instead of executing default logic.
