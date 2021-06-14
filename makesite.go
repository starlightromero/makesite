package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"strings"
)

type Data struct {
	Content string
}

func main() {
	var filename string

	flag.StringVar(&filename, "f", "", "name of file to write to html")
	flag.StringVar(&filename, "file", "", "name of file to write to html")

	flag.StringVar(&filename, "d", "", "name of directory to get all txt files")
	flag.StringVar(&filename, "dir", "", "name of directory to get all txt files")

	flag.Parse()

	fileContent := readFile(filename)
	fileToWrite := strings.SplitN(filename, ".", 2)[0]

	writeTemplate("template.tmpl", fileToWrite, string(fileContent))
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
