---
Title: Structs
Id: 6071
---
## List fields of a struct
Using reflection we can list all fields of a struct.

@file enum_struct_fields_simple.go output

Using reflection we can only access values (`v.Interface{}`) of exported fields.

Exported fields are fields with names starting with upper case (`FirstName` and `Age` are exported, `lastName` is not).

Field is exported if `reflect.StructField.PkgPath == ""`.


## List fields of a struct recursively

Inspecting a struct is inherently recursive process.

You have to chase pointers and recurse into embedded structures.

In real programs inspecting structures using reflections would be recursive.

@file enum_struct_fields.go output

We can access not only