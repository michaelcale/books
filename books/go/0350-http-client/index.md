---
Title: HTTP Client
Id: 1422
---

## Syntax
- resp, err := http.Get(url) // Makes a HTTP GET request with the default HTTP client. A non-nil error is returned if the request fails.
- resp, err := http.Post(url, bodyType, body) // Makes a HTTP POST request with the default HTTP client. A non-nil error is returned if the request fails.
- resp, err := http.PostForm(url, values) // Makes a HTTP form POST request with the default HTTP client. A non-nil error is returned if the request fails.

## Parameters
| Parameter | Details |
| ------ | ------ |
|resp   | A response of type `*http.Response` to an HTTP request |
|err   | An `error`. If not nil, it represents an error that occured when the function was called. |
|url   | A URL of type `string` to make a HTTP request to. |
|bodyType   | The MIME type of type `string` of the body payload of a POST request. |
|body   | An `io.Reader` (implements `Read()`) which will be read from until an error is reached to be submitted as the body payload of a POST request.|
|values   | A key-value map of type `url.Values`. The underlying type is a `map[string][]string`.|

## Remarks
It is important to `defer resp.Body.Close()` after every HTTP request that does not return a non-nil error, else resources will be leaked.

