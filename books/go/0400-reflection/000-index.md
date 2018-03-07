---
Title: Reflection
Id: 1854
---
Go is a statically typed language. In most cases the type of a variable is known at compilation time.

One exception is interface type, especially empty interface `interface{}`.

Empty interface is a dynamic type, similar to `Object` in Java or C#.

At compilation time we can't tell if the underlying value of interface type is an `int` or a `string`.

Package [`reflect`](https://golang.org/pkg/reflect/) in standard library allows us to work with such dynamic values at runtime. We can:
* inspect the type of dynamic value
* enumerate fields of a struct
* set the value
* create new values at runtime

Related language-level functionality for inspecting type of an interface value at runtime is a [type switch](a-14736) and a [type assertion](a-25362).

@file basic_reflect.go output

Basics of reflections are:
* start with value of empty interface `interface{}` type
* use `reflect.ValueOf(v interface{})` to get `reflect.Value` which represents information about the value
* use `reflect.Value` to check the type of value, test if the value is `nil`, set the value

## Uses for reflection

### Serialization

Reflection makes it possible to implement generic JSON serialization/deserialization.

For generic JSON serialization we can enumerate fields of arbitrary structres, read their fields and create corresponding JSON string.

For generic JSON deserialization, we can enumerate fields of arbitrary structures and set them based on parsed JSON data.

The same applies for other serialization formats like XML, YAML or Protocol Buffers.

Reflection makes it possible to define a generic API for SQL databases because we can convert arbitrary structures to a format that SQL database understands and put data recieved from SQL database into arbitrary structures.

### Extending templating language with Go functions

Thanks to ability to call arbitrary functions at runtime we can define custom function for templates in `text/template`. We register Go functions with templating engine.

The engine can then call those functions at runtime, when executing a template.

<!-- TODO: link to article showing how to define functions -->

### Writing interpreters tightly integrated with Go

Thanks to reflection's ability to call arbitrary functions at runtime, a JavaScript interpreter can be easily extended with additional functions written in Go.

<!-- TODO: examples golua, otto -->
