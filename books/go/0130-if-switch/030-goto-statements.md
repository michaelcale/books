---
Title: Goto statements
Id: 18834
Score: 3
---
A `goto` statement transfers control to the statement with the corresponding label within the same function.
Executing the `goto` statement must not cause any variables to come into scope that were not already in scope at the point of the `goto`.

For example see the standard library source code: https://golang.org/src/math/gamma.go :

```go
for x < 0 {
    if x > -1e-09 {
        goto small
    }
    z = z / x
    x = x + 1
}
for x < 2 {
    if x < 1e-09 {
        goto small
    }
    z = z / x
    x = x + 1
}

if x == 2 {
    return z
}

x = x - 2
p = (((((x*_gamP[0]+_gamP[1])*x+_gamP[2])*x+_gamP[3])*x+_gamP[4])*x+_gamP[5])*x + _gamP[6]
q = ((((((x*_gamQ[0]+_gamQ[1])*x+_gamQ[2])*x+_gamQ[3])*x+_gamQ[4])*x+_gamQ[5])*x+_gamQ[6])*x + _gamQ[7]
return z * p / q

small:
if x == 0 {
    return Inf(1)
}
return z / ((1 + Euler*x) * x)
```
