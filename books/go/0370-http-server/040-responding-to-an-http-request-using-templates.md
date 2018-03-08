---
Title: Responding to an HTTP Request using Templates
Id: 6822
Score: 1
---
Responses can be written to a `http.ResponseWriter` using templates in Go. This proves as a handy tool if you wish to create dynamic pages.

(To learn how Templates work in Go, please visit the [Go Templates Documentation](1402) page.)

Continuing with a simple example to utilise the `html/template` to respond to an HTTP Request:

```go
package main

import(
    "html/template"
    "net/http"
    "log"
)

func main(){
    http.HandleFunc("/",WelcomeHandler)
    http.ListenAndServe(":8080",nil)
}

type User struct{
    Name string
    nationality string //unexported field.
}

func check(err error){
    if err != nil{
        log.Fatal(err)
    }
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET"{
        t,err := template.ParseFiles("welcomeform.html")
        check(err)
        t.Execute(w,nil)
    }else{
        r.ParseForm()
        myUser := User{}
        myUser.Name = r.Form.Get("entered_name")
        myUser.nationality = r.Form.Get("entered_nationality")
        t, err := template.ParseFiles("welcomeresponse.html")
        check(err)
        t.Execute(w,myUser)
    }
}
```

Where, the contents of

1) `welcomeform.html` are:

```html
<head>
    <title> Help us greet you </title>
</head>
<body>
    <form method="POST" action="/">
        Enter Name: <input type="text" name="entered_name">
        Enter Nationality: <input type="text" name="entered_nationality">
        <input type="submit" value="Greet me!">
    </form>
</body>
```

1) `welcomeresponse.html` are:

```html
<head>
    <title> Greetings, {{.Name}} </title>
</head>
<body>
    Greetings, {{.Name}}.<br>
    We know you are a {{.nationality}}!
</body>
```

Note:

1) Make sure that the `.html` files are in the correct directory.

2) When `http://localhost:8080/` can be visited after starting the server.

3) As it can be seen after submitting the form, the *unexported* nationality field of the struct could not be parsed by the template package, as expected.
