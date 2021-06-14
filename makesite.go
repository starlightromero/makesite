package main

import (
	"html/template"
	"log"
	"os"
)

type Data struct {
	Content string
}

func main() {
	fileContent := readFile("first-post.txt")
	writeTemplate("template.tmpl", "first-post", string(fileContent))
}

func readFile(file string) []byte {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func writeTemplate(tmpl, filename, fileContent string) {
	data := Data{fileContent}

	htmlFile, osErr := os.Create(filename + ".html")
	if osErr != nil {
		log.Fatal(osErr)
	}

	t := template.Must(template.ParseFiles(tmpl))
	execErr := t.Execute(htmlFile, data)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
