package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	day := now.Weekday()
	hour := now.Hour()

	shour := ""
	switch {
		case hour < 12 : shour = "Morning"
		case hour < 17 : shour = "Afternoon"
		default : shour = "Evening"
	}

	dtype := ""
	switch day {
		case time.Saturday: fallthrough
		case time.Sunday: dtype = "WeekEnd"
		default: dtype = "Working Day"
	}
	
	fmt.Printf("%v %v %v is a %v\n", day, shour, now, dtype)
}

