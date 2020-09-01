package main

import (
	"fmt"

	"github.com/pluralsight/customtype/organization"
)

func main() {
	//var p organization.Identifiable = organization.NewPerson("James", "Wilson")
	p := organization.NewPerson("James", "Wilson", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))
	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.Country())
	println(p.FullName())
}
