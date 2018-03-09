---
Title: Normal Handler Function
Id: 297
Score: 0
SOId: 28935
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
