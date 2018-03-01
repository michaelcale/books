---
Title: HTTP Client
Id: 1422
---

Package 'net/http' in Go standard library provides functionality to make HTTP network requests.

In the examples we use [httpbin.org](http://httpbin.org/) which is a clever service that can return specific HTTP responses, which is useful for demonstrating various aspects of HTTP protocol.

## Basic HTTP GET

For simplicity this example uses `http.Get()`. In real programs you should use custom client with a timeout as described below.

@file http_get.go output noplayground

This shows how to make HTTP GET request for a URL (HTML page in this case).

I use `uri` as variable name because there is `net/url` package which means using more natural `url` will lead to naming conflicts when importing `net/url` in the same file.

When there is no error, `http.Get()` returns `*http.Response` with notable fields:
* `Body` is `io.Reader` containing content of the response. If the URL was HTML page, this is HTML data. If the URL was PNG image, this is PNG data
* `StatusCode` is an `int` describing HTTP status code. 200 means OK. 404 means Not Found etc.
* `Header` contains response headers. Its type is `map[string][]string` because per HTTP spec it's ok to have multiple headers with the same name

When no error is returned, it's important to `resp.Body.Close()` or you will leak resources.

Trying to access page that doesn't exist on the server returns a response wiht `StatusCode` 404 (Not Found).

Trying to access non-existent host will fail, in which case response is `nil`.

## HTTP GET using custom client

`http.Get()` is just a wrapper that delegates all work `http.DefaultClient`, which is a package variable of type `*http.Client`.

It's best not to use `http.Get` because default client doesn't have a timeout, which means it'll wait forever connecting to slow or buggy or malicious servers.

@file http_get_custom_client.go output allow_error noplayground

As shown, creating and using custom `http.Client` is easy.

In this example we set a very short time-out to demonstrate that exceeding this timeout cancels the connection.

In real programs we would use longer timeout, like 15 seconds (the exact timeout depends on particulars of your code).

## Basic HTTP POST

@file http_post.go output noplayground

The simplest way to do POST is using `http.Client.Post(url string, contentType string, body io.Reader)`.

In this example we send raw text. Most of the time the server expects body to be in url-encoded format.

## Basic HTTP HEAD

Like HTTP GET but use `http.Client.Head(uri string)` method.

