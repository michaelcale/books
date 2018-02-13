---
Title: Delete map element
Id: 2489
---
The [`delete`](https://golang.org/pkg/builtin/#delete) built-in function removes the element with the specified key from a map.

    people := map[string]int{"john": 30, "jane": 29}
    fmt.Println(people) // map[john:30 jane:29]

    delete(people, "john")
    fmt.Println(people) // map[jane:29]

If the `map` is `nil` or there is no such element, `delete` has no effect.

    people := map[string]int{"john": 30, "jane": 29}
    fmt.Println(people) // map[john:30 jane:29]

    delete(people, "notfound")
    fmt.Println(people) // map[john:30 jane:29]

    var something map[string]int
    delete(something, "notfound") // no-op
