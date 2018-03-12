---
Title: Primitive types
Id: 225
SOId: 6070
---

Let's see what kind of operations we can do on primitive types like `int` or `string`.

## Get the type

@file get_type.go output sha1:23f2449e0fa7db2890118c495f1d112f8fb35e18 goplayground:E6qpfwZiTEr

In real code you would handle [all the types](229) you care about.

## Get the value

@file get_value.go output allow_error sha1:aebd85272f71d9d6bcbadfca98feb0e7b4cd7ba3 goplayground:KuRi0TBk_5x

To minimize API surface, `Int()` returns `int64` and works on all signed integer values (`int8`, `int16`, `int32`, `int64`).

`UInt()` methods returns `uint64` and works on every unsigned integer values (`uint8`, `uint16`, `uint32`, `uint64`).

Trying to get integer value from value of incompatible type (like `string`) will panic.

To avoid panic you can first check the type with `Kind()`.

All methods for retrieving the value:

* `Bool() bool`
* `Int() int64`
* `UInt() uint64`
* `Float() float64`
* `String() string`
* `Bytes() []byte`

## Set the value

@file set_value.go output sha1:ee2ad75af1ab400651f45ac566b99429e822781c goplayground:O50qUNYUkJ0

As `setInt` and `setStructField` show, you can only change values if you start with a pointer to the value.

Since `reflect.ValueOf()` creates a `reflect.Value` that represents a pointer to a value, you need to use `Elem()` to get `reflect.Value` that represents the value itself. You can then call `SetInt()` to set the value.

`setStructPtrField` shows how we can grab a reference to field value by it's position in the struct.

Trying to set value of incompatible type will panic.

Methods that set values mirror those that read the values:

* `SetBool(v bool)`
* `SetInt(v int64)`
* `SetUInt(v uint64)`
* `SetFloat(v float64)`
* `SetString(v string)`
* `SetBytes(v []byte)`
