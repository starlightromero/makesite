package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Hello, world!")
	fileContents := readFile("first-post.txt")
	fmt.Print(string(fileContents))
}

func readFile(file string) []byte {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return fileContents
}
