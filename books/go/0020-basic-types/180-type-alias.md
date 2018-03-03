---
Title: Type aliases
Id: 901000v5
---
Type aliases were introduced in Go 1.9 to make [code refactoring](https://talks.golang.org/2016/refactor.article) easier.

Imagine you have package `foo` that exports type `Bar`.

You want to rename type `Bar` to `NewBar`.

Without type aliases you would have to change all the packages that use `foo.Bar` to use `foo.NewBar` at the same time you're renaming the type.

With type aliases you can split this into a 2 step process.

First update dependencies by introducing an alias for `foo.Bar` and replacing all uses of `foo.Bar` with the alias.

```go
import "foo"
type Bar = foo.Bar // Bar is now an alias of foo.Bar
```

Now you can rename `foo.Bar` to `foo.NewBar` and update the alias:
```go
import "foo"
type Bar = foo.NewBar
```

This is a much smaller change.

You can now gradually get rid of the alias and use the new `foo.NewBar` type directly.

The process sounds like a hassle, but in large code bases it can be a better approach than renaming everything at once.

It's tempting to use type aliases for other things, but you should resist.

Type aliases add a layer of indirection, which hurts readability of the code. There should be a good reason to use them.

To learn more, you can [read the spec](https://github.com/golang/proposal/blob/master/design/18130-type-alias.md).
