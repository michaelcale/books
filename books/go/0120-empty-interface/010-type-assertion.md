---
Title: Type assertion
Id: 95
Score: 0
SOId: 25362
---
At compile time, when you have a variable whose type is [interface](90) (including [empty interface](94)) you don't know what is the real, underlying type.

You can access underlying type at runtime using type assertion.

@file type_assertion.go output allow_error sha1:c97c73fb1d526862c1cda162f8df757ca6ba1a5e goplayground:knuECI96ypQ

Another way of accessing underlying type is with [type switch](96).

For completness, you can use short version of type switch: `v := iv.(int)` (vs. `v, ok := iv.(int)`).

The difference is that the short version will panic if `iv` is not of the asserted type:

@file type_assertion2.go output allow_error sha1:e56c2798f794acc37ea8fccd08bc294b8abfebbc goplayground:4nyHsjDu4i0

As a rule of thumb, you shouldn't try to discover underlying value of interface type as it pierces through an abstraction.
