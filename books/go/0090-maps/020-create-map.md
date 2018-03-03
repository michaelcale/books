---
Title: Create a map
Id: 2484
---
One can declare and initialize a map in a single statement using a [*composite literal*](https://golang.org/ref/spec#Composite_literals).

Using automatic type Short variable declaration:

    mapIntInt := map[int]int{10: 100, 20: 100, 30: 1000}
    mapIntString := map[int]string{10: "foo", 20: "bar", 30: "baz"}
    mapStringInt := map[string]int{"foo": 10, "bar": 20, "baz": 30}
    mapStringString := map[string]string{"foo": "one", "bar": "two", "baz": "three"}

The same code, but with Variable types:

    var mapIntInt = map[int]int{10: 100, 20: 100, 30: 1000}
    var mapIntString = map[int]string{10: "foo", 20: "bar", 30: "baz"}
    var mapStringInt = map[string]int{"foo": 10, "bar": 20, "baz": 30}
    var mapStringString = map[string]string{"foo": "one", "bar": "two", "baz": "three"}

You can also include your own structs in a map:

You can use custom types as the value:

    // Custom struct types
    type Person struct {
      FirstName, LastName string
    }

    var mapStringPerson = map[string]Person{
      "john": Person{"John", "Doe"},
      "jane": Person{"Jane", "Doe"}}
    mapStringPerson := map[string]Person{
      "john": Person{"John", "Doe"},
      "jane": Person{"Jane", "Doe"}}

Your struct can also be the _key_ to the map:

    type RouteHit struct {
        Domain string
        Route  string
    }

    var hitMap = map[RouteHit]int{
      RouteHit{"example.com","/home"}: 1,
      RouteHit{"example.com","/help"}: 2}
    hitMap := map[RouteHit]int{
      RouteHit{"example.com","/home"}: 1,
      RouteHit{"example.com","/help"}: 2}


You can create an empty map simply by not entering any value within the brackets `{}`.

    mapIntInt := map[int]int{}
    mapIntString := map[int]string{}
    mapStringInt := map[string]int{}
    mapStringString := map[string]string{}
    mapStringPerson := map[string]Person{}

You can create and use a map directly, without the need to assign it to a variable. However, you will have to specify both the declaration and the content.

    // using a map as argument for fmt.Println()
    fmt.Println(map[string]string{
      "FirstName": "John",
      "LastName": "Doe",
      "Age": "30"})

    // equivalent to
    data := map[string]string{
      "FirstName": "John",
      "LastName": "Doe",
      "Age": "30"}
    fmt.Println(data)
