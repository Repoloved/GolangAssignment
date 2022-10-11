package main

import (
	"fmt"
	"time"
	"math"
	"strings"
)

// date time format layout
const TextDate = "2 January 2006"

// ! NOTE: Count Days
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

// ! NOTE: Present Employee
func GetSalary(rangeDay int, data [][]string) map[string]string {
	return nil // TODO: replace this
}

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi helper untuk mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
// ! Convert int to Format Rupiah
func FormatRupiah(number int) string {
	hitung := rupiah.FormatIntToRp(number)
	return hitung
}

func GetSalaryOverview(dateRange int, data [][]string) map[string]string {
	var money
	var name

	dateRange = GetDayDifference()
	data = make(map[string] string)
	data[name] = dateRange * money
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
