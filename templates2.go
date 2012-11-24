package main

import (
    "os"
    "text/template"
    "fmt"
)

type Person struct {
    Name string
    nonExportedAgeField string
}

func main() {
    p := Person{Name: "Mary", nonExportedAgeField: "44"}
    t := template.New("nonexported template demo")
    t, _ = t.Parse("Hello {{.Name}}! Age is {{.nonExportedAgeField}}.")
    err := t.Execute(os.Stdout, p)
    if err != nil {
        fmt.Println("There was an error:", err)
    }
}
