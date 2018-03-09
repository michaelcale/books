---
Title: Composition and embedding
Id: 84
SOId: 1256
---
Composition provides an alternative to inheritance. A struct may include another type by name in its declaration:

```go
type Request struct {
    Resource string
}

type AuthenticatedRequest struct {
    Request
    Username, Password string
}
```

In the example above, `AuthenticatedRequest` will contain four public members: `Resource`, `Request`, `Username`, and `Password`.

Composite structs can be instantiated and used the same way as normal structs:

```go
func main() {
    ar := new(AuthenticatedRequest)
    ar.Resource = "example.com/request"
    ar.Username = "bob"
    ar.Password = "P@ssw0rd"
    fmt.Printf("%#v", ar)
}
```

## Embedding

In the previous example, `Request` is an embedded field. Composition can also be achieved by embedding a different type. This is useful, for example, to decorate a Struct with more functionality. For example, continuing with the Resource example, we want a function that formats the content of the Resource field to prefix it with `http://` or `https://`. We have two options: create the new methods on AuthenticatedRequest or **embed** it from a different struct:

```go
type ResourceFormatter struct {}

func(r *ResourceFormatter) FormatHTTP(resource string) string {
    return fmt.Sprintf("http://%s", resource)
}
func(r *ResourceFormatter) FormatHTTPS(resource string) string {
    return fmt.Sprintf("https://%s", resource)
}

type AuthenticatedRequest struct {
    Request
    Username, Password string
    ResourceFormatter
}

```

And now the main function could do the following:

```go
func main() {
    ar := new(AuthenticatedRequest)
    ar.Resource = "www.example.com/request"
    ar.Username = "bob"
    ar.Password = "P@ssw0rd"

    println(ar.FormatHTTP(ar.Resource))
    println(ar.FormatHTTPS(ar.Resource))

    fmt.Printf("%#v", ar)
}
```

Look that the `AuthenticatedRequest` that has a `ResourceFormatter` embedded struct.

**But** the downside of it is that you cannot access objects outside of your composition. So `ResourceFormatter` cannot access members from `AuthenticatedRequest`.
