---
Title: Go <-> JSON type mappings
Id: 188
SOId: g32600tb
---

This is the mapping between Go types and JSON types:

| JSON Type | Go Concrete Type |
| ------ | ------ |
| boolean   | bool   |
| numbers   | float64 or int   |
| string   | string   |
| array | slice |
| dictionary | struct |
| null   | nil   |

Here's how they look in practice:

@file type_mapping.go output sha1:98f4f0fb3439e97bc8e73268a73dbb378a9990c9 goplayground:F36DkvViH06