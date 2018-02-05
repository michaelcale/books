---
Title: Characters and runes
Search: char, rune
Id: 9010002e
---
Go has 2 type characters:
* `byte` is 1 byte value, an alias for `uint8` type
* `rune` is 4 byte Unicode code-point, an alias for `int32` type

[Zero value](a-6069) of a `byte` or `rune` is 0.

## Iterate over string using bytes

```go
s := "str"
for i := 0; i < len(s); i++ {
    c := s[i]
    fmt.Printf("Byte at index %d is %c (0x%x)\n", i, c, c)
}
```
**Output:**
```text
Byte at index 0 is s (0x73)
Byte at index 1 is t (0x74)
Byte at index 2 is r (0x72)
```

## Iterate over string using runes

```go
s := "日本語"
for i, runeChar := range s {
    fmt.Printf("Rune at byte position %d is %#U\n", i, runeChar)
}
```
**Output**:
```text
Rune at byte position 0 is U+65E5 '日'
Rune at byte position 3 is U+672C '本'
Rune at byte position 6 is U+8A9E '語'
```
