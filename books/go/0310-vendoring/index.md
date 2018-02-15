---
Title: Vendoring
Id: 978
---

## Remarks
Vendoring is a method of ensuring that all of your 3rd party packages that you use in your Go project are consistent for everyone who develops for your application.

When your Go package imports another package, the compiler normally checks `$GOPATH/src/` for the path of the imported project. However if your package contains a folder named `vendor`, the compiler will check in that folder *first*. This means that you can import other parties packages inside your own code repository, without having to modify their code.

Vendoring is a standard feature in Go 1.6 and up. In Go 1.5, you need to set the environment variable of `GO15VENDOREXPERIMENT=1` to enable vendoring.
