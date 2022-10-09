package main

import (
	"fmt"
	"time"
	"math"
	"strings"
)

// date time format layout
const TextDate = "2 January 2006"

func GetDayDifference(date string)int {
	splitmin := strings.Split(date, " - ")
	splitspace := strings.Split(date, " ")
	start_date := splitmin[0] + " " + splitspace[5]
	end_date := splitmin[1]
	date1, _ := time.Parse(TextDate, start_date)
	date2, _ := time.Parse(TextDate, end_date)
	days := int(math.Ceil(date2.Sub(date1).Hours() / 24))
	return days
}

func main() {
	res := GetDayDifference("21 February - 23 February 2021")
    fmt.Println(res)
}