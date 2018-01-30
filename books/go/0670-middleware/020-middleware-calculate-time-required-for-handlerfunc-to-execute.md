Title: Middleware Calculate time required for handlerFunc to execute
Id: 28936
Score: 0
Body:
    // logger middlerware that logs time taken to process each request
    func Logger(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            startTime := time.Now()
            h.ServeHttp(w,r)
            endTime := time.Since(startTime)
            log.Printf("%s %d %v", r.URL, r.Method, endTime)
        })
    }
    
    func loginHandler(w http.ResponseWriter, r *http.Request) {
                // Steps to login
    }
    
    
    func main() {
        http.HandleFunc("/login", Logger(loginHandler))
        http.ListenAndServe(":8080", nil)
    }
        
|======|
