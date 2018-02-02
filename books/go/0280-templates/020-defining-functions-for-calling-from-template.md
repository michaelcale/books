Title: Defining functions for calling from template
Id: 13535
Score: 2
Body:
 

    package main
    
    import (
        "fmt"
        "net/http"
        "os"
        "text/template"
    )
    
    var requestTemplate string = `
    {{range $i, $url := .URLs}}
    {{ $url }} {{(status_code $url)}}
    {{ end }}`
    
    type Requests struct {
        URLs []string
    }
    
    func main() {
        var fns = template.FuncMap{
            "status_code": func(x string) int {
                resp, err := http.Head(x)
                if err != nil {
                    return -1
                }
                return resp.StatusCode
            },
        }
    
        req := new(Requests)
        req.URLs = []string{"http://godoc.org", "http://stackoverflow.com", "http://linux.org"}
    
        tmpl := template.Must(template.New("getBatch").Funcs(fns).Parse(requestTemplate))
        err := tmpl.Execute(os.Stdout, req)
        if err != nil {
            fmt.Println(err)
        }
    }

Here we use our defined function `status_code` to get status code of web page right from template.

Output:

    http://godoc.org 200
    
    http://stackoverflow.com 200
    
    http://linux.org 200
|======|
