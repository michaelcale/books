---
Title: Iterating the elements of a map
Id: 2486
Score: 7
---

    import fmt

    people := map[string]int{
      "john": 30,
      "jane": 29,
      "mark": 11,
    }

    for key, value := range people {
      fmt.Println("Name:", key, "Age:", value)
    }

Note that when iterating over a map with a range loop, [the iteration order is not specified](https://blog.golang.org/go-maps-in-action) and is not guaranteed to be the same from one iteration to the next.

You can also discard either the keys or the values of the map, if you are looking to just [grab keys](http://stackoverflow.com/documentation/go/732/maps/2487/iterating-the-keys-of-a-map#t=201607220900510314955) or just grab values.
