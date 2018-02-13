---
Title: goto statements
Id: 18834
Score: 3
---
A `goto` statement transfers control to the statement with the corresponding label within the same function.

Executing the `goto` statement must not cause any variables to come into scope that were not already in scope at the point of the `goto`.

@file goto.go output
