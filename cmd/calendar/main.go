package main

import (
	"fmt"
	"time"
)

func printCalendar(year int, month time.Month) {
	// Create a date for the first day of the month
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	// Print the month and year
	fmt.Printf("%s %d\n", date.Format("January"), year)

	// Print the days of the week
	fmt.Println(" Mo Tu We Th Fr Sa Su")

	// Print the calendar for the month
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			day := date.AddDate(0, 0, i*7+j-int(date.Weekday()))
			if day.Month() != date.Month() {
				fmt.Print("   ")
			} else {
				fmt.Printf("%3d ", day.Day())
			}
		}
		fmt.Println()
	}
}

func main() {
	now := time.Now()
	printCalendar(now.Year(), now.Month())
}
