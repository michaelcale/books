---
Title: Configuring JSON struct fields
Id: 14157
Score: 1
---
Consider the following example:

```go
type Company struct {
    Name     string
    Location string
}
```

## Hide/skip certain fields

To export `Revenue` and `Sales`, but hide them from encoding/decoding, use `json:"-"` or rename the variable to begin with a lowercase letter. Note that this prevents the variable from being visible outside the package.

```go
type Company struct {
    Name     string `json:"name"`
    Location string `json:"location"`
    Revenue  int    `json:"-"`
    sales    int
}
```

## Ignore empty fields

To prevent `Location` from being included in the JSON when it is set to its zero value, add `,omitempty` to the `json` tag.

```go
type Company struct {
    Name     string `json:"name"`
    Location string `json:"location,omitempty"`
}
```

