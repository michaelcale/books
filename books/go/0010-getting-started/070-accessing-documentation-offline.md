---
Title: Accessing documentation offline
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

For quick reference check from command-line, eg for fmt.Print:

```sh
godoc cmd/fmt Print
# or
go doc fmt Print
```

General help is also available from command-line:

```sh
go help [command]
```
