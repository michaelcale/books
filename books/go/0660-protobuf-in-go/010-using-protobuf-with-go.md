---
Title: Using Protobuf with Go
Id: 29998
Score: 2
---
The message you want to serialize and send that you can include into a file **test.proto**, containing

```go
package example;

enum FOO { X = 17; };

message Test {
    required string label = 1;
    optional int32 type = 2 [default=77];
    repeated int64 reps = 3;
    optional group OptionalGroup = 4 {
    required string RequiredField = 5;
    }
}
```

To compile the protocol buffer definition, run protoc with the --go_out parameter set to the directory you want to output the Go code to.

```sh
protoc --go_out=. *.proto
```

To create and play with a Test object from the example package,

```go
package main

import (
    "log"

    "github.com/golang/protobuf/proto"
    "path/to/example"
)

func main() {
    test := &example.Test {
        Label: proto.String("hello"),
        Type:  proto.Int32(17),
        Reps:  []int64{1, 2, 3},
        Optionalgroup: &example.Test_OptionalGroup {
            RequiredField: proto.String("good bye"),
        },
    }
    data, err := proto.Marshal(test)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }
    newTest := &example.Test{}
    err = proto.Unmarshal(data, newTest)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }
    // Now test and newTest contain the same data.
    if test.GetLabel() != newTest.GetLabel() {
        log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
    }
    // etc.
}
```

To pass extra parameters to the plugin, use a comma-separated parameter list separated from the output directory by a colon:

```sh
protoc --go_out=plugins=grpc,import_path=mypackage:. *.proto
```
