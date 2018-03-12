---
Title: Creating a token using a custom claims type
Id: 310
Score: 0
SOId: 31137
---
The `StandardClaim` is embedded in the custom type to allow for easy encoding, parsing and validation of standard claims.

```go
tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

type MyCustomClaims struct {
    Foo string `json:"foo"`
    jwt.StandardClaims
}

// sample token is expired.  override time so it parses as valid
at(time.Unix(0, 0), func() {
    token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("AllYourBase"), nil
    })

    if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
        fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
    } else {
        fmt.Println(err)
    }
})
```

Output:

```text
bar 15000
```

(From the [documentation](https://godoc.org/github.com/dgrijalva/jwt-go#ex-ParseWithClaims--CustomClaimsType), courtesy of Dave Grijalva.)
