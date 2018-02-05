---
Title: Type aliases
Id: 901000v5
---
Type alias was introduced in Go 1.9 to make [code refactoring](https://talks.golang.org/2016/refactor.article) easier.

Imagine you have package `foo` that exports type `Bar`.

You want to rename type `Bar` to `NewBar`.

Without type aliases you have to update all packages that use `foo.Bar` to use `foo.NewBar` at the same time when you're renaming the type.

With type aliases you can split this into 2 step process.

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

This is much smaller change.

You can now gradually get rid of the alias and use the new `foo.NewBar` type directly.

The process sounds like a hassle, but in large code bases it can be better approach than renaming everything at once.

It's tempting to use type aliases for other things, but you should resist.

Type alias adds a layer of indirection, which hurts readability of the code. There should be a good reason to pay for that.

To learn more, you can [read the spec](https://github.com/golang/proposal/blob/master/design/18130-type-alias.md).
