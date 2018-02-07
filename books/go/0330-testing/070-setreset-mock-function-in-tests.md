---
Title: Set/Reset Mock Function In Tests
Id: 12204
Score: 1
---
This example shows how to mock out a function call that is irrelevant to our unit test, and then use the `defer` statement to re-assign the mocked function call back to its original function.

```go
var validate = validateDTD

// ParseXML parses b for XML elements and values, and returns them as a map of
// string key/value pairs.
func ParseXML(b []byte) (map[string]string, error) {
    // we don't care about validating against DTD in our unit test
    if err := validate(b); err != nil {
        return err
    }

    // code to parse b etc.
}

func validateDTD(b []byte) error {
    // get the DTD from some external storage, use it to validate b etc.
}
```

In our unit test,

```go
func TestParseXML(t *testing.T) {
    // assign the original validate function to a variable.
    originalValidate = validate
    // use the mockValidate function in this test.
    validate = mockValidate
    // defer the re-assignment back to the original validate function.
    defer func() {
       validate = originalValidate
    }()

    var input []byte
    actual, err := ParseXML(input)
    // assertion etc.
}

func mockValidate(b []byte) error {
    return nil // always return nil since we don't care
}
```
