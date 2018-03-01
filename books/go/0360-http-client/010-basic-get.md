---
Title: HTTP POST
Id: 4643
---

HTTP POST sends binary data to the server.

How this data is interpreted by the server depends on the value of `Content-Type` header.

## HTTP POST with url-encoded data

HTTP POST is most often used to submit filled forms data in HTML page to the server.

Form data is is a dictionary of key/value pairs where key is a string and value is array of strings (most often array with a single element).

Form key/value pairs are most often sent as url-encoded data with `Content-Type` of `application/x-www-form-urlencoded`.

@file http_post_url_encoded.go output noplayground
