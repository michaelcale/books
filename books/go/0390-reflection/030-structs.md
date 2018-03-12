---
Title: Structs
Id: 227
SOId: 6071
---

## List fields of a struct

Using reflection we can list all fields of a struct.

@file enum_struct_fields_simple.go output sha1:9538eecc07c20cd6c0ce463e54670f82dc6e6061 goplayground:21KyxJi0sNK

Using reflection we can only access values (`v.Interface{}`) of exported fields.

Exported fields are fields with names starting with upper case (`FirstName` and `Age` are exported, `lastName` is not).

Field is exported if `reflect.StructField.PkgPath == ""`.

## List fields of a struct recursively

Inspecting a struct is inherently recursive process.

You have to chase pointers and recurse into embedded structures.

In real programs inspecting structures using reflections would be recursive.

@file enum_struct_fields.go output sha1:15a48e2a1c07d976ad58570626543420ecfef224 goplayground:tmplO4Fkru7

We can access not only
