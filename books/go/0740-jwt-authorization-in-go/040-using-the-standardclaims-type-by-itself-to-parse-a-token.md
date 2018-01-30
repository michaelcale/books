Title: Using the StandardClaims type by itself to parse a token
Id: 31139
Score: 0
Body:
The `StandardClaims` type is designed to be embedded into your custom types to provide standard validation features. You can use it alone, but there's no way to retrieve other fields after parsing. See the custom claims example for intended usage.

    mySigningKey := []byte("AllYourBase")
    
    // Create the Claims
    claims := &jwt.StandardClaims{
        ExpiresAt: 15000,
        Issuer:    "test",
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    ss, err := token.SignedString(mySigningKey)
    fmt.Printf("%v %v", ss, err)

Output:

    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.QsODzZu3lUZMVdhbO76u3Jv02iYCvEHcYVUI1kOWEU0 <nil>

(From the [documentation](https://godoc.org/github.com/dgrijalva/jwt-go#ex-NewWithClaims--StandardClaims), courtesy of Dave Grijalva.)
|======|
