Title: Encoding/Decoding using Go structs
Id: 22028
Score: 0
Body:
Let's assume we have the following `struct` that defines a `City` type:

    type City struct {  
        Name string  
        Temperature int  
    }

We can encode/decode City values using the [`encoding/json`](https://golang.org/pkg/encoding/json/) package.

First of all, we need to use the Go metadata to tell the encoder the correspondence between the struct fields and the JSON keys.

    type City struct {  
        Name string `json:"name"`  
        Temperature int `json:"temp"`  
        // IMPORTANT: only exported fields will be encoded/decoded  
        // Any field starting with a lower letter will be ignored  
    }  

To keep this example simple, we'll declare an explicit correspondence between the fields and the keys. However, you can use several variants of the `json:` metadata [as explained in the docs](https://golang.org/pkg/encoding/json/#Marshal).

**IMPORTANT:** **Only [exported fields](http://stackoverflow.com/documentation/go/374/structs/1255/exported-vs-unexported-fields-private-vs-public#t=201607220824119778568) (fields with capital name) will be serialized/deserialized.** For example, if you name the field _**t**emperature_ it will be ignored even if you set the `json` metadata.

## Encoding

To encode a `City` struct, use `json.Marshal` as in the basic example:

    // data to encode  
    city := City{Name: "Rome", Temperature: 30}  
     
    // encode the data  
    bytes, err := json.Marshal(city)  
    if err != nil {  
        panic(err)  
    }  
     
    fmt.Println(string(bytes))  
    // {"name":"Rome","temp":30} 

[Playground](https://play.golang.org/p/KlziJIDWPW)

## Decoding

To decode a `City` struct, use `json.Unmarshal` as in the basic example:

    // data to decode  
    bytes := []byte(`{"name":"Rome","temp":30}`)  
     
    // initialize the container for the decoded data  
    var city City  
     
    // decode the data  
    // notice the use of &city to pass the pointer to city  
    if err := json.Unmarshal(bytes, &city); err != nil {  
        panic(err)  
    }  
     
    fmt.Println(city)  
    // {Rome 30} 

[Playground](https://play.golang.org/p/VHS28E-234)
|======|
