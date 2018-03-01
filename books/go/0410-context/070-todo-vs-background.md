---
Title: context.TODO() vs. context.Background()
Id: 901000a2
---
You can create new, empty context using [`context.TODO()`](https://golang.org/pkg/context/#TODO) and [`context.Background()`](https://golang.org/pkg/context/#Background).

What's the difference?

In terms of functionality: none. They are exactly the same value, bit-by-bit.

The difference is in intent.

The docs describe `context.TODO()` as:

> TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter). TODO is recognized by static analysis tools that determine whether Contexts are propagated correctly in a program.

The docs describe `context.Background()` as:

> Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.

Frankly, I'm not sure what they are trying to say.

I guess `context.TODO()` is meant to be used if you expect that at some point in the future you'll no longer need to create context there, either because it'll be passed from the outside or that there will be a more specific way to create it.

If you can't decide, don't sweat it. In practice they behave the same way so pick whichever.
