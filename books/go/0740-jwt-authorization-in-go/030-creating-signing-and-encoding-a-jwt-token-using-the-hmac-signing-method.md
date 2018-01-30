Title: Creating, signing, and encoding a JWT token using the HMAC signing method
Id: 31138
Score: 0
Body:
    // Create a new token object, specifying signing method and the claims
    // you would like it to contain.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "foo": "bar",
        "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
    })
    
    // Sign and get the complete encoded token as a string using the secret
    tokenString, err := token.SignedString(hmacSampleSecret)
    
    fmt.Println(tokenString, err)

Output:

    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU <nil>

(From the [documentation](https://godoc.org/github.com/dgrijalva/jwt-go#ex-New--Hmac), courtesy of Dave Grijalva.)
|======|
