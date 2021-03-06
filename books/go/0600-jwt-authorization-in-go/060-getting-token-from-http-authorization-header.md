---
Title: Getting token from HTTP Authorization header
Id: 314
Score: 0
SOId: 31141
---
```go
type contextKey string

const (
    // JWTTokenContextKey holds the key used to store a JWT Token in the
    // context.
    JWTTokenContextKey contextKey = "JWTToken"

    // JWTClaimsContextKey holds the key used to store the JWT Claims in the
    // context.
    JWTClaimsContextKey contextKey = "JWTClaims"
)

// ToHTTPContext moves JWT token from request header to context.
func ToHTTPContext() http.RequestFunc {
    return func(ctx context.Context, r *stdhttp.Request) context.Context {
        token, ok := extractTokenFromAuthHeader(r.Header.Get("Authorization"))
        if !ok {
            return ctx
        }

        return context.WithValue(ctx, JWTTokenContextKey, token)
    }
}
```

(From [go-kit/kit](https://github.com/go-kit/kit/blob/master/auth/jwt/transport.go), courtesy of Peter Bourgon)
