package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println("header")
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println("content")
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println("footer")
	s1.Execute(os.Stdout, nil)
}
