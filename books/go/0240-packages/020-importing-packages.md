---
Title: Importing packages
Id: 22421
Score: 1
---

You can import a single package with the statement:

```go
import "path/to/package"
```

or group multiple imports together:

```go
import (
    "path/to/package1"
    "path/to/package2"
)
```

This will look in the corresponding `import` paths inside of the `$GOPATH` for `.go` files and lets you access exported names through `packagename.AnyExportedName`.

You can also access local packages inside of the current folder by prefacing packages with `./`. In a project with a structure like this:

```test
project
├── src
│   ├── package1
│   │   └── file1.go
│   └── package2
│       └── file2.go
└── main.go
```

You could call this in `main.go` in order to import the code in `file1.go` and `file2.go`:

```go
import (
    "./src/package1"
    "./src/package2"
)
```

Since package-names can collide in different libraries you may want to alias one package to a new name. You can do this by prefixing your import-statement with the name you want to use.

```go
import (
    "fmt" //fmt from the standardlibrary
    tfmt "some/thirdparty/fmt" //fmt from some other library
)
```

This allows you to access the former `fmt` package using `fmt.*` and the latter `fmt` package using `tfmt.*`.

You can also import the package into the own namespace, so that you can refer to the exported names without the `package.` prefix using a single dot as alias:

```go
import (
    . "fmt"
)
```

Above example imports `fmt` into the global namespace and lets you call, for example, `Printf`.

If you import a package but don't use any of it's exported names, the Go compiler will print an error-message. To circumvent this, you can set the alias to the underscore:

```go
import (
    _ "fmt"
)
```

This can be useful if you don't access this package directly but need it's `init` functions to run.

**Note:**

As the package names are based on the folder structure, any changes in the folder names & import references (including case sensitivity) will cause a compile time error "case-insensitive import collision" in Linux & OS-X, which is difficult to trace and fix (the error message is kinda cryptic for mere mortals as it tries to convey the opposite - that, the comparison failed due to case sensitivity).

ex: "path/to/Package1" vs "path/to/package1"
