---
Title: Use golang/dep
Id: 290
Score: 2
SOId: 29860
---
[golang/dep](https://github.com/golang/dep) is a prototype dependency management tool. Soon to be an official versioning tool. Current status **Alpha**.

## Usage

Get the tool via
```sh
$ go get -u github.com/golang/dep/...
```

Typical usage on a new repo might be

```sh
$ dep init
$ dep ensure -update
```

To update a dependency to a new version, you might run

```sh
$ dep ensure github.com/pkg/errors@^0.8.0
```

Note that the manifest and lock file formats **have now been finalized**. These will remain compatible even as the tool changes.
