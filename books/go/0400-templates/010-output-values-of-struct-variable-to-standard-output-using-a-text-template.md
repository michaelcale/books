Title: Output values of struct variable to Standard Output using a text template
Id: 4575
Score: 2
Body:
    package main
    
    import (
        "log"
        "text/template"
        "os"
    )
    
    type Person struct{
        MyName string
        MyAge int
    }
    
    var myTempContents string= `
    This person's name is : {{.MyName}}
    And he is {{.MyAge}} years old.
    `
    
    func main() {
        t,err := template.New("myTemp").Parse(myTempContents)
        if err != nil{
            log.Fatal(err)
        }
        myPersonSlice := []Person{ {"John Doe",41},{"Peter Parker",17} }
        for _,myPerson := range myPersonSlice{
            t.Execute(os.Stdout,myPerson)
        }
    }

[Playground](https://play.golang.org/p/HwaxzuwO7A)
|======|
