---
Title: Custom XML marshaling
Id: 335
---

## Writing custom XML marshalling

Sometimes a type doesn't have an obvious mapping to JSON.

How would you serialize `time.Time`? There are so many possibilities.

Go provides a default XML mapping for `time.Time`. We can implement custom marshaller for user-defined types like structs.

For existing types we can define a new (but compatible) type.

Here's a custom marshalling and unmarshaling for `time.Time` that only serializes year/month/date part:

@file custom_marshal.go output

Notice that receiver type of `UnmashalXML` is a pointer to the type.

This is necessary for changes to persist outside the function itself.


## Custom marshaling behind the scenes

How does custom marshaling works?

Package XML defines 2 interfaces: [`Marshaler`](https://golang.org/pkg/encoding/xml/#Marshaler) and [`Unmarshaler`](https://golang.org/pkg/encoding/xml/#Unmarshaler).

```go
type Marshaler interface {
    MarshalXML(e *Encoder, start StartElement) error
}

type Unmarshaler interface {
    UnmarshalXML(d *Decoder, start StartElement) error
}
```

By implementing those functions we make our type conform to `Marshaler` or `Unmarshaler` interface.

XML encoder / decoder checks if the value being encoded conforms to those interfaces and will call those functions instead of executing default logic.
