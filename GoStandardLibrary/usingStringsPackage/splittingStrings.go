package main

import (
	"fmt"
	"strings"
)

func main() {
	ourString := "this is a string!"
	stringCollection := strings.Split(ourString, " ")
	for i := range stringCollection {
		fmt.Println(stringCollection[i])
	}

	anotherString := "this is| another string!"
	stringCollection = strings.SplitAfter(anotherString, "|")
	for i := range stringCollection {
		fmt.Println(stringCollection[i])
	}
}
