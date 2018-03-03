---
Title: Access documentation offline
Id: 24998
---
For full documentation, run the command:

```sh
godoc -http=:<port-number>
```

For a tour of Go (highly recommended for beginners in the language):

```sh
go tool tour
```

The two commands above will start web-servers with documentation similar to what is found online [here](https://golang.org/doc/) and [here](https://tour.golang.org/) respectively.

You can also use `godoc` for quick reference. For example, to see documentation for fmt.Print:

```sh
godoc cmd/fmt Print
# or
go doc fmt Print
```

General help is also available from the command-line:

```sh
go help [command]
```
