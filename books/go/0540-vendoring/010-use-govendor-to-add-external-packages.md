---
Title: Use govendor to add external packages
Id: 288
Score: 7
SOId: 3183
---
[Govendor](https://github.com/kardianos/govendor) is a tool that is used to import 3rd party packages into your code repository in a way that is compatible with golang's vendoring.

Say for example that you are using a 3rd party package `bosun.org/slog`:

```go
package main

import "bosun.org/slog"

func main() {
    slog.Infof("Hello World")
}
```

Your directory structure might look like:

```text
$GOPATH/src/
├── github.com/me/helloworld/
|   ├── hello.go
├── bosun.org/slog/
|   ├── ... (slog files)
```

However someone who clones `github.com/me/helloworld` may not have a `$GOPATH/src/bosun.org/slog/` folder, causing _their_ build to fail due to missing packages.

Running the following command at your command prompt will grab all the external packages from your Go package and package the required bits into a vendor folder:

    govendor add +e

This instructs govendor to add all of the external packages into your current repository.

Your application's directory structure would now look like:

```text
$GOPATH/src/
├── github.com/me/helloworld/
|   ├── vendor/
|   |   ├── bosun.org/slog/
|   |   |   ├── ... (slog files)
|   ├── hello.go
```

and those who clone your repository will also grab the required 3rd party packages.

