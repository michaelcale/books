Title: Go Get
Id: 28395
Score: 1
Body:
`go get` downloads the packages named by the import paths, along with their
dependencies. It then installs the named packages, like 'go install'. Get also accepts build flags to control the installation.

> go get github.com/maknahar/phonecountry

When checking out a new package, get creates the target directory
`$GOPATH/src/<import-path>`. If the GOPATH contains multiple entries,
get uses the first one. Similarly, it will install compiled binaries in `$GOPATH/bin`. 

When checking out or updating a package, get looks for a branch or tag
that matches the locally installed version of Go. The most important
rule is that if the local installation is running version "go1", get
searches for a branch or tag named "go1". If no such version exists it
retrieves the most recent version of the package.

When using `go get`, the `-d` flag causes it to download but not install the given package. The `-u` flag will allow it to update the package and its dependencies.

Get never checks out or updates code stored in vendor directories.
|======|
