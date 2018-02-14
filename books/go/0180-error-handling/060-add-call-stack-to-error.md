---
Title: Add callstack to error messages
Id: k1k100ea
---

If you come from languages like Java, C# or Python, you might be used to the fact that exceptions include call stack for the location that created the exception.

Collecting callstacks is expensive and Go doesn't add callstack to errors.

Callstacks are useful in debugging. If you're ok with additional cost, you can use package [github.com/pkg/errors](https://godoc.org/github.com/pkg/errors) to augment errors with callstack.

@file add_callstack.go output

As you can see top part of the callstack includes Go runtime code.
