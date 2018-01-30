Title: Maps
Id: 732
Introduction:
Maps are data types used for storing unordered key-value pairs, so that looking up the value associated to a given key is very efficient. Keys are unique. The underlying data structure grows as needed to accommodate new elements, so the programmer does not need to worry about memory management. They are similar to what other languages call hash tables, dictionaries, or associative arrays.
|======|
Syntax:
- var mapName map[KeyType]ValueType  // declare a Map
- var mapName = map[KeyType]ValueType{}  // declare and assign an empty Map
- var mapName = map[KeyType]ValueType{key1: val1, key2: val2}  // declare and assign a Map
- mapName := make(map[KeyType]ValueType) // declare and initialize default size map
- mapName := make(map[KeyType]ValueType, length) // declare and initialize *length* size map
- mapName := map[KeyType]ValueType{}  // auto-declare and assign an empty Map with :=
- mapName := map[KeyType]ValueType{key1: value1, key2: value2}  // auto-declare and assign a Map with :=
- value := mapName[key]  // Get value by key
- value, hasKey := mapName[key]  // Get value by key, 'hasKey' is 'true' if key exists in map
- mapName[key] = value  // Set value by key

|======|
Remarks:
Go provides a built-in `map` type that implements a _hash table_. Maps are Go's built-in associative data type (also called _hashes_ or _dictionaries_ in other languages).

|======|
