Title: Iterating the values of a map
Id: 7117
Score: 3
Body:
    people := map[string]int{
      "john": 30,
      "jane": 29,
      "mark": 11,
    }
    
    for _, value := range people {
      fmt.Println("Age:", value)
    }

Note that when iterating over a map with a range loop, [the iteration order is not specified](https://blog.golang.org/go-maps-in-action) and is not guaranteed to be the same from one iteration to the next. 

|======|
