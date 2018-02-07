---
Title: Iterating the elements of a map
Id: 2486
Score: 7
---

```go
import fmt

people := map[string]int{
  "john": 30,
  "jane": 29,
  "mark": 11,
}

for key, value := range people {
  fmt.Println("Name:", key, "Age:", value)
}
```

Iteration order is not specified. In fact, Go on purpose randomizes the order of iteration so that programmers don't write buggy code by relying on specific order.

You can also discard either the keys or the values of the map, if you are looking to just [grab keys](a-2487) or just grab values.
