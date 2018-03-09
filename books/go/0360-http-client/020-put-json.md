---
Title: PUT request of JSON object
Id: 201
SOId: 27703
---
`http.Client` doesn't have a convenience method for doing PUT requests, so we construct a `http.Request` object and use `http.Client.Do(req *http.Request)` to perform that request.

@file http_put.go output noplayground

