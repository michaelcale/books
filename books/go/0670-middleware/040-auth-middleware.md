Title: Auth Middleware
Id: 28938
Score: 0
Body:
    func Authenticate(h http.Handler) http.Handler {
        return CustomHandlerFunc(func(w *http.ResponseWriter, r *http.Request) {
            // extract params from req
            // post params | headers etc
            if CheckAuth(params) {
                log.Println("Auth Pass")
                // pass control to next middleware in chain or handler func
                h.ServeHTTP(w, r)
            } else {
                log.Println("Auth Fail")
                // Responsd Auth Fail
            }
        })
    }
|======|
