---
Title: Blank identifier
Id: 29103
---
To help avoid mistakes, the Go compiler doesn't allow unused variables.

However, there are some situations when you don't need to use a value stored in a variable.

In those cases, you use a "blank identifier" `_` to assign and discard the assigned value.

A blank identifier can be assigned a value of any type, and is most commonly used in functions that return multiple values.

**Multiple Return Values**

@file blank_identifier.go output sha1:a9c22c5a4db87bfdaac2f74c3d56fa108f94fe84 goplayground:q5i3MnL_qBO

**Using `range`**

@file blank_identifier_2.go output sha1:13781d48412168b4f7f8fbb1c9c0b373bd023c36 goplayground:c9m512EHc2p

