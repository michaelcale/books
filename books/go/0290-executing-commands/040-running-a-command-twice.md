---
Title: Running a Command twice
Id: 23356
Score: 0
---
> A Cmd cannot be reused after calling its Run, Output or CombinedOutput methods

Running a command twice *will **not** work*:

```go
cmd := exec.Command("xte", "key XF86AudioPlay")
_ := cmd.Run() // Play audio key press
// .. do something else
err := cmd.Run() // Pause audio key press, fails
```

> Error: exec: already started

Rather, one must use **two separate** `exec.Command`. You might also need some delay between commands.

```go
cmd := exec.Command("xte", "key XF86AudioPlay")
_ := cmd.Run() // Play audio key press
// .. wait a moment
cmd := exec.Command("xte", "key XF86AudioPlay")
_ := cmd.Run() // Pause audio key press
```
