---
Title: Using a map as a set
Id: 14398
Score: 4
---

Some languages have a native structure for sets. To make a set in Go, it's best practice to use a map from the value type of the set to an empty struct (`map[Type]struct{}`).

For example, with strings:

@file maps_as_sets.go output
