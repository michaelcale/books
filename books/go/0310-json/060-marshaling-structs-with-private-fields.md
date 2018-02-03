Title: Marshaling structs with private fields
Id: 14194
Score: 0
Body:
As a good developer you have created following struct with both exported and unexported fields:

```
type MyStruct struct {
    uuid string    
    Name string
}
```

Example in Playground:
https://play.golang.org/p/Zk94Il2ANZ

Now you want to `Marshal()` this struct into valid JSON for storage in something like etcd. However, since `uuid` in not exported, the `json.Marshal()` skips it. What to do? Use an anonymous struct and the `json.MarshalJSON()` interface! Here's an example:

```
type MyStruct struct {
    uuid string    
    Name string
}

func (m MyStruct) MarshalJSON() ([]byte, error {
    j, err := json.Marshal(struct {
        Uuid string
        Name string
    }{
        Uuid: m.uuid,
        Name: m.Name,
    })
    if err != nil {
           return nil, err
    }
    return j, nil
}
```

Example in Playground:
https://play.golang.org/p/Bv2k9GgbzE


|======|
