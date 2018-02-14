---
Title: Panic
Id: 17009
---
A panic halts normal execution flow and exits the current function.

Any deferred calls will then be executed before control is passed to the next function higher on the stack.

Each stack's function will exit and run deferred calls until the panic is handled using a deferred `recover()`, or until the panic reaches `main()` and terminates the program.

If this occurs, the argument provided to panic and a stack trace will be printed to `stderr`.

@file index.go output allow_error

`panic` accepts any type as its parameter.
