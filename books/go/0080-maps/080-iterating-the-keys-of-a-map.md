---
Title: Iterating the keys of a map
Id: 2487
Score: 6
---
    people := map[string]int{
      "john": 30,
      "jane": 29,
      "mark": 11,
    }

    for key, _ := range people {
      fmt.Println("Name:", key)
    }

If you are just looking for the keys, since they are the first value, you can simply drop the underscore:

    for key := range people {
      fmt.Println("Name:", key)
    }

Note that when iterating over a map with a range loop, [the iteration order is not specified](https://blog.golang.org/go-maps-in-action) and is not guaranteed to be the same from one iteration to the next.
