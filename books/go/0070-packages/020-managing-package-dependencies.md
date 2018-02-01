Title: Managing package dependencies
Id: 6680
Score: 1
Body:
A common way to download Go dependencies is by using the `go get <package>` command, which will save the package into the global/shared `$GOPATH/src` directory.  This means that a single version of each package will be linked into each project that includes it as a dependency. This also means that when a new developers deploys your project, they will `go get` the latest version of each dependency.

However you can keep the build environment consistent, by attaching all the dependencies of a project into the `vendor/` directory. Keeping vendored dependencies committed along with your project's repository allows you to do per-project dependency versioning, and provide a consistent environment for your build.

This is what your project's structure will look like:

    $GOPATH/src/
    ├── github.com/username/project/
    |   ├── main.go 
    |   ├── vendor/
    |   |   ├── github.com/pkg/errors
    |   |   ├── github.com/gorilla/mux
|======|
