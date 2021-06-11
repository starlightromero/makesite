package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileContents := readFile("first-post.txt")
	fmt.Print(string(fileContents))
}

func readFile(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
