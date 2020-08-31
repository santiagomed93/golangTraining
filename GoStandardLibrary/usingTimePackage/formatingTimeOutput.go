package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
)

func main() {
	currentTime := time.Now()
	startDate := time.Date(2018, 07, 04, 9, 00, 00, 00, time.UTC)

	fmt.Println(currentTime.Format(layoutISO))
	fmt.Println(startDate.Format(layoutUS))
	fmt.Println(currentTime.Format(layoutEU))
	fmt.Println(startDate.Format(time.ANSIC))
	fmt.Println(currentTime.Format(time.RFC1123))

	Year := currentTime.Year()
	Month := currentTime.Month()
	Day := currentTime.Day()
	fmt.Printf("Today is %d/%d/%d\n", Month, Day, Year)
}
