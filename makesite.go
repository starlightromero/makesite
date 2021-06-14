package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Data struct {
	Content string
}

func main() {
	var filename string
	var directory string
	var fileCount int

	flag.StringVar(&filename, "f", "", "name of file to write to html")
	flag.StringVar(&filename, "file", "", "name of file to write to html")

	flag.StringVar(&directory, "d", "", "name of directory to get all txt files to write to html")
	flag.StringVar(&directory, "dir", "", "name of directory to get all txt files to write to html")

	flag.Parse()

	if directory != "" {
		fileCount += writeAllFilesToHTML(directory)
	}

	if filename != "" {
		fileContent := readFile(filename)
		fileToWrite := stripExt(filename)

		fileCount += writeToHTML("template.tmpl", fileToWrite, string(fileContent))
	}

	if len(os.Args) < 2 {
		fmt.Println("file or dir flag is required")
		os.Exit(1)
	}

	boldGreen := color.New(color.FgGreen, color.Bold)
	white := color.New(color.FgWhite)
	boldWhite := color.New(color.FgWhite, color.Bold)

	boldGreen.Print("Success! ")
	white.Print("Generated ")
	boldWhite.Print(fileCount)
	white.Print(" pages.")

}

func readFile(file string) []byte {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func writeToHTML(tmpl, filename, fileContent string) int {
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

	return 1
}

func writeAllFilesToHTML(directory string) int {
	var fileCount int

	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if isTxt(file.Name()) {
			filename := stripExt(file.Name())
			fileContent := readFile(file.Name())
			writeToHTML("template.tmpl", filename, string(fileContent))
			fileCount += 1
		}
	}

	return fileCount
}

func isTxt(filename string) bool {
	fileExt := filename[len(filename)-4:]
	return fileExt == ".txt"
}

func stripExt(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}
