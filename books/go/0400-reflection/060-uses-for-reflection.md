---
Title: Uses for reflection
Id: 23400
---
## Serialization to JSON, XML, SQL, protobufs etc.

Reflection makes it possible to implement generic JSON serialization/deserialization.

For generic JSON serialization we can enumerate fields of arbitrary structres, read their fields and create corresponding JSON string.

For generic JSON deserialization, we can enumerate fields of arbitrary structures and set them based on parsed JSON data.

The same applies for other serialization formats like XML, YAML or Protocol Buffers.

Reflection makes it possible to define a generic API for SQL databases because we can convert arbitrary structures to a format that SQL database understands and put data recieved from SQL database into arbitrary structures.

## Extending templating language with Go functions

Thanks to ability to call arbitrary functions at runtime we can define custom function for templates in `text/template`. We register Go functions with templating engine.

The engine can then call those functions at runtime, when executing a template.

<!-- TODO: link to article showing how to define functions -->

## Writing interpreters tightly integrated with Go

Thanks to reflection's ability to call arbitrary functions at runtime, a JavaScript interpreter can be easily extended with additional functions written in Go.

<!-- TODO: examples golua, otto -->
