package main

import (
	"fmt"
	"time"
	"math"
	"strings"
)

// date time format layout
const TextDate = "2 January 2006"

func GetDayDifference(date string) int {
	splitmin := strings.Split(date, " - ")
	splitspace := strings.Split(date, " ")
	start_date := splitmin[0] + " " + splitspace[5]
	end_date := splitmin[1]
	date1, _ := time.Parse(TextDate, start_date)
	date2, _ := time.Parse(TextDate, end_date)
	days := int(math.Ceil(date2.Sub(date1).Hours() / 24))
	return days
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	return nil // TODO: replace this
}

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi helper untuk mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	return "" // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	return nil // TODO: replace this
}

func main() {
	res := GetSalaryOverview("21 February - 23 February 2021", [][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	fmt.Println(res)
}
