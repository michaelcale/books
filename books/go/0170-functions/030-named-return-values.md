Title: Named Return Values
Id: 1253
Body:
Return values can be assigned to a local variable. An empty `return` statement can then be used to return their current values. This is known as *"naked"* return. Naked return statements should be used only in short functions as they harm readability in longer functions:

```go
func Inverse(v float32) (reciprocal float32) {
    if v == 0 {
        return
    }
    reciprocal = 1 / v
    return
}
```
[play it on playground](https://play.golang.org/p/dS_bGmP6W0)

```go
//A function can also return multiple values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```
[play it on playground](https://play.golang.org/p/upOAwpOaue)

Two important things must be noted:

- The parenthesis around the return values are **mandatory**.
- An empty `return` statement must always be provided.
