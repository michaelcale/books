---
Title: if action
Id: 219
SOId: 13458
---

To conditionally render parts of the template, use `if` action:

@file if.go output

## false values

Templating uses "truthy" logic for deciding what values are true or false in the context of `if` action:

@file if2.go output

## Avoid printing empty slices

Truthy logic is useful when we want to show different text if a list of items is empty:

@file if_empty_slice.go output
