---
Title: Use map as a set
Id: 14398
---

Some languages have a native structure for sets. To make a set in Go, it's best practice to use a map from the value type of the set to an empty struct (`map[Type]struct{}`).

For example, with strings:

@file maps_as_sets.go output sha1:9356d0423461567e7128b60aea266d356ef7e33b goplayground:CCCOYYSl2R0
