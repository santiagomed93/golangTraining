package main

import (
	"fmt"

	"github.com/pluralsight/customtype/media"
)

func main() {
	var myFav media.Catalogable = &media.Movie{}
	fmt.Printf("el tipo es %T\n", myFav)
	myFav.NewMovie("Top Gun", media.PG, 32.8)

	fmt.Printf("My favorite movie is %s\n", myFav.GetTitle())
	fmt.Printf("It was rated %v\n", myFav.GetRating())
	fmt.Printf("It made %.1f in the box office\n", myFav.GetBoxOffice())

	myFav.SetTitle("Dumb and Dumber") // Create a new copy of the struct with a new title - performance issue
	fmt.Printf("My favorite movie is now %s\n", myFav.GetTitle())
}
