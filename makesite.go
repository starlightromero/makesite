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
	var directory string

	flag.StringVar(&filename, "f", "", "name of file to write to html")
	flag.StringVar(&filename, "file", "", "name of file to write to html")

	flag.StringVar(&directory, "d", "", "name of directory to get all txt files to write to html")
	flag.StringVar(&directory, "dir", "", "name of directory to get all txt files to write to html")

	flag.Parse()

	if directory != "" {
		printAllTxtFiles(directory)
	}

	if filename != "" {
		fileContent := readFile(filename)
		fileToWrite := stripExt(filename)

		writeToHTML("template.tmpl", fileToWrite, string(fileContent))
	}
}

func readFile(file string) []byte {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func writeToHTML(tmpl, filename, fileContent string) {
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

func printAllTxtFiles(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if isTxt(file.Name()) {
			filename := stripExt(file.Name())
			fileContent := readFile(file.Name())
			writeToHTML("template.tmpl", filename, string(fileContent))
		}
	}
}

func isTxt(filename string) bool {
	fileExt := filename[len(filename)-4:]
	return fileExt == ".txt"
}

func stripExt(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}
