---
Title: Time out request with a context
Id: 12209
Score: 5
---
## 1.7+

Timing out an HTTP request with a context can be accomplished with only the standard library (not the subrepos) in 1.7+:

```go
import (
    "context"
    "net/http"
    "time"
)

req, err := http.NewRequest("GET", `https://example.net`, nil)
ctx, _ := context.WithTimeout(context.TODO(), 200 * time.Milliseconds)
resp, err := http.DefaultClient.Do(req.WithContext(ctx))
// Be sure to handle errors.
defer resp.Body.Close()
```

## Before 1.7

```go
import (
    "net/http"
    "time"

    "golang.org/x/net/context"
    "golang.org/x/net/context/ctxhttp"
)

ctx, err := context.WithTimeout(context.TODO(), 200 * time.Milliseconds)
resp, err := ctxhttp.Get(ctx, http.DefaultClient, "https://www.example.net")
// Be sure to handle errors.
defer resp.Body.Close()
```

## Further Reading

For more information on the `context` read [chapter on context](a-2743).
