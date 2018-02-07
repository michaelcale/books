---
Title: Defer
Id: 2795
---
## Introduction
A `defer` statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.

### Syntax
- defer someFunc(args)
- defer func(){
    //code goes here
  }()

## Remarks
Defer works by injecting a new stack frame (the called function after the `defer` keyword) into the call stack below the currently executing function. This means that defer is guaranteed to run as long as the stack will be unwound (if your program crashes or gets a `SIGKILL`, defer will not execute).
