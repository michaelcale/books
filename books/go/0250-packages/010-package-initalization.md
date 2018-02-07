---
Title: Package initalization
Id: 6074
Score: 1
---

Package can have `init` methods which are run **only once** before main.

```go
package usefull

func init() {
    // init code
}
```

If you just want to run the package initialization without referencing anything from it use the following import expression.

```go
    import _ "usefull"
```
