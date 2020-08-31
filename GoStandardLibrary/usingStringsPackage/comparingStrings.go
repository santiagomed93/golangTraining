package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "this is a string"
	string2 := "this is a string"

	if string1 == string2 {
		fmt.Println("The strings are the same")
	}

	if strings.Compare(string1, string2) == 0 {
		fmt.Println("The strings are the same")
	}
}
