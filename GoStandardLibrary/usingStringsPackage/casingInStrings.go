package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Never trust a programmer who carries a screwddriver\n"
	fmt.Println("Before: " + sampleString)

	strLowerCase := strings.ToLower(sampleString)
	fmt.Println("Lower case: " + strLowerCase)

	strUpperCase := strings.ToUpper(sampleString)
	fmt.Println("Upper case: " + strUpperCase)

	strTitleCase := strings.Title(sampleString)
	fmt.Println("Title case: " + strTitleCase)

	fmt.Println("Title case using function: " + properTitle(sampleString))
}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "
	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
