package main

import (
    "text/template"
    "os"
    "fmt"
    )

func main() {
    fmt.Println("Load a set of templates with {{define}} clauses and execute:")
    s1, _ := template.ParseFiles("t1.tmpl", "t2.tmpl") //create a set of templates from many files.
    //Note that t1.tmpl is the file with contents "{{define "t_ab"}}a b{{template "t_cd"}}e f {{end}}"
    //Note that t2.tmpl is the file with contents "{{define "t_cd"}} c d {{end}}"


    s1.ExecuteTemplate(os.Stdout, "t_cd", nil) //just printing of c d
    fmt.Println()
    s1.ExecuteTemplate(os.Stdout, "t_ab", nil) //execute t_ab which will include t_cd
    fmt.Println()
    s1.Execute(os.Stdout, nil) //since templates in this data structure are named, there is no default template and so it prints nothing
}
