package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"longshuai", 23}
	tmpl, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	_ = tmpl.Execute(os.Stdout, p)
}
