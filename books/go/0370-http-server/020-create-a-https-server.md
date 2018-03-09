---
Title: Create a HTTPS Server
Id: 208
Score: 7
SOId: 3248
---

## Generate a certificate
In order to run a HTTPS server, a certificate is necessary. Generating a self-signed certificate with `openssl` is done by executing this command:

```sh
openssl req -x509 -newkey rsa:4096 -sha256 -nodes -keyout key.pem -out cert.pem -subj "/CN=example.com" -days 3650`
```

The parameters are:

 - `req` Use the certificate request tool
 - `x509` Creates a self-signed certificate
 - `newkey rsa:4096` Creates a new key and certificate by using the RSA algorithms with `4096` bit key length
 - `sha256` Forces the SHA256 hashing algorithms which major browsers consider as secure (at the year 2017)
 - `nodes` Disables the password protection for the private key. Without this parameter, your server had to ask you for the password each time its starts.
 - `keyout` Names the file where to write the key
 - `out` Names the file where to write the certificate
 - `subj` Defines the domain name for which this certificate is valid
 - `days` Fow how many days should this certificate valid? `3650` are approx. 10 years.

Note: A self-signed certificate could be used e.g. for internal projects, debugging, testing, etc. Any browser out there will mention, that this certificate is not safe. In order to avoid this, the certificate must signed by a certification authority. Mostly, this is not available for free. One exception is the "Let's Encrypt" movement: https://letsencrypt.org

## The necessary Go code

You can handle configure TLS for the server with the following code. `cert.pem` and `key.pem` are your SSL certificate and key, which where generated with the above command.

```go
package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })

    log.Fatal(http.ListenAndServeTLS(":443","cert.pem","key.pem", nil))
}
```

