---
Title: Package initalization
Id: 6074
---
Package can have one or more `init` methods.

Those methods are run **only once** before `main` function is executed.

```go
package foo

func init() {
    // init code
}
```

If you just want to run the package initialization without referencing anything from it use the following import expression.

```go
import _ "foo"
```

Package initialization function can be used to create initial state needed by code in the package.

Avoid temptation of doing too much in `init` function. Such implicit logic negatively impacts other people's ability to understand the code.
