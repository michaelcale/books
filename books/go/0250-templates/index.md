Title: Templates
Id: 1402
Syntax:
- t, err := template.Parse(`{{.MyName .MyAge}}`)

- t.Execute(os.Stdout,struct{MyValue,MyAge string}{"John Doe","40.1"})
|======|
Remarks:
Golang provides packages like:

 1. `text/template`

 2. `html/template`

to implement data-driven templates for generating textual and HTML outputs.
|======|
