package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Data struct {
	Content string
}

func main() {
	var filename string
	var directory string
	var fileCount int
	var fileSize float64

	start := time.Now()

	flag.StringVar(&filename, "f", "", "name of file to write to html")
	flag.StringVar(&filename, "file", "", "name of file to write to html")

	flag.StringVar(&directory, "d", "", "name of directory to get all txt files to write to html")
	flag.StringVar(&directory, "dir", "", "name of directory to get all txt files to write to html")

	flag.Parse()

	if directory != "" {
		fileCount, fileSize = writeAllFilesToHTML(directory)
	}

	if filename != "" {
		fileContent := readFile(filename)
		fileCount, fileSize = writeToHTML("template.tmpl", filename, string(fileContent))
	}

	if len(os.Args) < 2 {
		fmt.Println("file or dir flag is required")
		os.Exit(1)
	}

	boldGreen := color.New(color.FgGreen, color.Bold)
	white := color.New(color.FgWhite)
	boldWhite := color.New(color.FgWhite, color.Bold)

	end := time.Now()
	elapsed := end.Sub(start)
	milliseconds := elapsed.Seconds() / 1000.0

	boldGreen.Print("Success! ")
	white.Print("Generated ")
	boldWhite.Print(fileCount)
	white.Printf(" pages (%.1fkB total) in %.2f milliseconds.", fileSize, milliseconds)

}

func readFile(file string) []byte {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func writeToHTML(tmpl, filePath, fileContent string) (int, float64) {
	data := Data{fileContent}

	htmlFilePath := strings.SplitN(filePath, ".", 2)[0] + ".html"
	htmlFile, osErr := os.Create(htmlFilePath)
	if osErr != nil {
		log.Fatal(osErr)
	}

	t := template.Must(template.ParseFiles(tmpl))
	execErr := t.Execute(htmlFile, data)
	if execErr != nil {
		log.Fatal(execErr)
	}

	return 1, getFileSize(htmlFilePath)
}

func writeAllFilesToHTML(directory string) (int, float64) {
	var fileCount int
	var fileSize float64

	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := directory + "/" + file.Name()

		info, statErr := os.Stat(path)
		if statErr != nil {
			log.Fatal(statErr)
		}

		if info.IsDir() {
			writeAllFilesToHTML(path)
		} else {
			if isTxt(file.Name()) {
				fileContent := readFile(path)
				count, size := writeToHTML("template.tmpl", path, string(fileContent))
				fileCount += count
				fileSize += size
			}
		}
	}

	return fileCount, fileSize
}

func isTxt(filename string) bool {
	fileExt := filename[len(filename)-4:]
	return fileExt == ".txt"
}

func getFileSize(filename string) float64 {
	file, openErr := os.Open(filename)
	if openErr != nil {
		log.Fatal(openErr)
	}

	size, seekErr := file.Seek(0, 2)
	if seekErr != nil {
		log.Fatal(seekErr)
	}

	return float64(size) / 1024.0
}
