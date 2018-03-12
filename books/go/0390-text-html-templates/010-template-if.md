---
Title: if action
Id: 219
SOId: 13458
---

To conditionally render parts of the template, use `if` action:

@file if.go output sha1:01dc196565b5e4fdfb8c6ce923ffc753b6239acc goplayground:mcVhMbPtV9T

## false values

Templating uses "truthy" logic for deciding what values are true or false in the context of `if` action:

@file if2.go output sha1:ce4fdad0dbaa80360959af98f1fd362fd03ed030 goplayground:s8GCT7H34EL

## Avoid printing empty slices

Truthy logic is useful when we want to show different text if a list of items is empty:

@file if_empty_slice.go output sha1:ea01ab225c326934713d5cf909bce693fd8a057c goplayground:7xSoiyH0QDJ
