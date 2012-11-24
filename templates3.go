package main

import (
    "text/template"
    "fmt"
    )

func main() {
    tOk := template.New("first")
    template.Must(tOk.Parse(" some static text /* and a comment */")) //a valid template, so no panic with Must
    fmt.Println("The first one parsed OK.")

    template.Must(template.New("second").Parse("some static text {{ .Name }}")) 
    fmt.Println("The second one parsed OK.")

    fmt.Println("The next one ought to fail.")
    tErr := template.New("check parse error with Must")
    template.Must(tErr.Parse(" some static text {{ .Name }")) // due to unmatched brace, there should be a panic here
}
