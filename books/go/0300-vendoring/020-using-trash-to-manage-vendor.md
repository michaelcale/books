---
Title: Using trash to manage ./vendor
Id: 26497
Score: 2
---
[`trash`](https://github.com/rancher/trash) is a minimalistic vendoring tool that you configure with `vendor.conf` file. This example is for `trash` itself:

```text
# package
github.com/rancher/trash

github.com/Sirupsen/logrus                      v0.10.0
github.com/urfave/cli                           v1.18.0
github.com/cloudfoundry-incubator/candiedyaml   99c3df8  https://github.com/imikushin/candiedyaml.git
github.com/stretchr/testify                     v1.1.3
github.com/davecgh/go-spew                      5215b55
github.com/pmezard/go-difflib                   792786c
golang.org/x/sys                                a408501
```

The first non-comment line is the package we're managing ./vendor for (note: this can be literally any package in your project, not just the root one).

Commented lines begin with `#`.

Each non-empty and non-comment line lists one dependency. Only the "root" package of the dependency needs to be listed.

After the package name goes the version (commit, tag or branch) and optionally the package repository URL (by default, it's inferred from the package name).

To populate your ./vendor dir, you need to have `vendor.conf` file in the current dir and just run:

```sh
$ trash
```

Trash will clone the vendored libraries into `~/.trash-cache` (by default), checkout requested versions, copy the files into `./vendor` dir and **prune non-imported packages and test files**. This last step keeps your ./vendor lean and mean and helps save space in your project repo.

Note: as of v0.2.5 trash is available for Linux and macOS, and only supports git to retrieve packages, as git's the most popular one, but we're working on adding all the others that `go get` supports.
