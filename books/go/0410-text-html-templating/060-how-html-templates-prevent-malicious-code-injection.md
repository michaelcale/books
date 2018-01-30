Title: How HTML templates prevent malicious code injection
Id: 13462
Score: 0
Body:
First, here's what can happen when `text/template` is used for HTML. Note Harry's `FirstName` property).

    package main

    import (
        "fmt"
        "html/template"
        "os"
    )

    type Person struct {
        FirstName string
        LastName  string
        Street    string
        City      string
        State     string
        Zip       string
        AvatarUrl string
    }

    func main() {
        const (
            letter = `<html><body><table>
    <tr><th></th><th>Name</th><th>Address</th></tr>
    {{range .}}
    <tr>
    <td><img src="{{.AvatarUrl}}"></td>
    <td>{{.FirstName}} {{.LastName}}</td>
    <td>{{.Street}}, {{.City}}, {{.State}} {{.Zip}}</td>
    </tr>
    {{end}}
    </table></body></html>`
        )

        tmpl, err := template.New("letter").Parse(letter)
        if err != nil {
            fmt.Println(err.Error())
        }

        harry := Person{
            FirstName: `Harry<script>alert("You've been hacked!")</script>`,
            LastName:  "Jones",
            Street:    "1234 Main St.",
            City:      "Springfield",
            State:     "IL",
            Zip:       "12345-6789",
            AvatarUrl: "harry.png",
        }

        jane := Person{
            FirstName: "Jane",
            LastName:  "Sherman",
            Street:    "8511 1st Ave.",
            City:      "Dayton",
            State:     "OH",
            Zip:       "18515-6261",
            AvatarUrl: "jane.png",
        }

        tmpl.Execute(os.Stdout, []Person{harry, jane})
    }

Results in:

    <html><body><table>
    <tr><th></th><th>Name</th><th>Address</th></tr>
    
    <tr>
    <td><img src="harry.png"></td>
    <td>Harry<script>alert("You've been hacked!")</script> Jones</td>
    <td>1234 Main St., Springfield, IL 12345-6789</td>
    </tr>
    
    <tr>
    <td><img src="jane.png"></td>
    <td>Jane Sherman</td>
    <td>8511 1st Ave., Dayton, OH 18515-6261</td>
    </tr>
    
    </table></body></html>

The above example, if accessed from a browser, would result in the script being executed an an alert being generated. If, instead, the `html/template` were imported instead of `text/template`, the script would be safely sanitized:

    <html><body><table>
    <tr><th></th><th>Name</th><th>Address</th></tr>
    
    <tr>
    <td><img src="harry.png"></td>
    <td>Harry&lt;script&gt;alert(&#34;You&#39;ve been hacked!&#34;)&lt;/script&gt; Jones</td>
    <td>1234 Main St., Springfield, IL 12345-6789</td>
    </tr>
    
    <tr>
    <td><img src="jane.png"></td>
    <td>Jane Sherman</td>
    <td>8511 1st Ave., Dayton, OH 18515-6261</td>
    </tr>
    
    </table></body></html>

The second result would look garbled when loaded in a browser, but would not result in a potentially malicious script executing.
|======|
