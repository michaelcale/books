---
Title: Normal Handler Function
Id: 28935
Score: 0
---
```go
func loginHandler(w http.ResponseWriter, r *http.Request) {
            // Steps to login
}

func main() {
    http.HandleFunc("/login", loginHandler)
    http.ListenAndServe(":8080", nil)
}
```
