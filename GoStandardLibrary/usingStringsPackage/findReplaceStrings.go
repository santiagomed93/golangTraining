package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "I really enjoy the Go language. It's one of my favorites"
	searchTerm := "Go"

	result := strings.Contains(sampleString, searchTerm)
	result2 := strings.HasPrefix(sampleString, "I")
	result3 := strings.HasSuffix(sampleString, "vorites")
	fmt.Printf("%v %v %v\n", result, result2, result3)

	ourString := "This is my string, There are many strings like it, but this one is mine"
	ourString = strings.Replace(ourString, "string", "compile", -1)
	fmt.Println(ourString)
}
