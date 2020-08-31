package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "      This is our Text      "
	newString := strings.TrimSpace(sampleString)
	newString2 := strings.TrimLeft(sampleString, " ")
	newString3 := strings.TrimRight(sampleString, " ")
	fmt.Printf("%q\n", newString)
	fmt.Printf("%q\n", newString2)
	fmt.Printf("%q\n", newString3)

	urlString := "https://www.something.com"
	domainName := strings.TrimPrefix(urlString, "https://")
	fmt.Println(domainName)
}
