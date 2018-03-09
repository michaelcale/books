---
Title: 3 ways of doing HTTP requests
Id: 205
SOId: 801000d0
---
Package `net/http` has a layered design where each layer is a convenience wrapper on top of a lower layer.

Each lower layer is more complex but offers more control.

Here's a recap of 3 ways of doing an HTTP GET request.

## Use `http.Get()` function

Using top-level `http.Get()` function is the simplest but not recommended due to lack of timeouts.

## Use `http.Client.Get()` method

* create `*http.Client` with `&http.Client{}`
* set appropriate `Timeout`
* use its `Get()` or `Post()` or `PostForm()` methods

## Use `http.Client.Do()` method

This allows the greatest control over the request.

* create `http.Client` and set apropriate `Timeout`
* create `*http.Request` with `http.NewRequest`
* set its `Method` to `"GET"`, `"POST"`, `"HEAD"`, `"PUT"` or `"DELETE"`
* set custom headers like `User-Agent`with `Header.Set(key, value string)`
* set body to be sent by setting `Body` of type `io.ReadCloser`
* set `TransferEncoding`
* send the request with `client.Do(req *http.Request)`
* set per-request timeout by creating a new request with context containing deadline `req = request.WithContext(ctx)`

This is the only way to do PUT or DELETE requests.
