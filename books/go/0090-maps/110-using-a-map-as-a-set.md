---
Title: Using a map as a set
Id: 14398
Score: 4
---
Some languages have a native structure for sets. To make a set in Go, it's best practice to use a map from the value type of the set to an empty struct (`map[Type]struct{}`).

For example, with strings:

    // To initialize a set of strings:
    greetings := map[string]struct{}{
        "hi":    {},
        "hello": {},
    }

    // To delete a value:
    delete(greetings, "hi")

    // To add a value:
    greetings["hey"] = struct{}{}

    // To check if a value is in the set:
    if _, ok := greetings["hey"]; ok {
        fmt.Println("hey is in greetings")
    }
