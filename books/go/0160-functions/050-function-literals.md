---
Title: Function literals
Id: 117
SOId: 1265
---

A **function literal** is a function declared inline, without a name.

Simplest example of a function literal:

@file function_literal.go output sha1:4fad807c897dd63cf0ab9991f707e92a779972e4 goplayground:BWL7hfMGKxh

Function literal with arguments:

@file function_literal2.go output sha1:b1287a42459723bad4a9817225aea6391154c84b goplayground:35zR7cUfvBt

Function literal closing over variable `str`:

@file function_literal3.go output sha1:26b13caaa692a3ccaa68c9d863c8bca0cdfd2728 goplayground:jD6pfOUE8vK

Function that closes over variables is called a [closure](118).

<!-- TODO: describe gotcha about argument evaluation -->
