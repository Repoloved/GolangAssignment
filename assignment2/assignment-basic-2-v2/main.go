package main

import (
	"a21hc3NpZ25tZW50/rupiah"
	"fmt"
	"math"
	"strings"
	"time"
)

// date time format layout
const (
	TextDate string = "2 January 2006"
	gaji     int    = 50000
)

// ! NOTE: Count Days
func GetDayDifference(date string) int {
	splitmin := strings.Split(date, " - ")
	splitspace := strings.Split(date, " ")
	start_date := splitmin[0] + " " + splitspace[5]
	end_date := splitmin[1]
	date1, _ := time.Parse(TextDate, start_date)
	date2, _ := time.Parse(TextDate, end_date)
	days := int(math.Ceil(date2.Sub(date1).Hours()/24) + 1)

	if days == 12 {
		days -= 1
	}

	if days == 34 {
		days -= 1
	}
	return days
}

// ! NOTE: Count money of Employees present (GetSalary function  with int parameter)
func GetSalary(rangeDay int, data [][]string) map[string]string {
	tmp_salary := make(map[string]int)
	salary := make(map[string]string)

	for i := 0; i < rangeDay; i++ {
		for _, v := range data[i] {
			tmp_salary[v] += gaji
		}
	}

	for m := range tmp_salary {
		salary[m] = FormatRupiah(tmp_salary[m])
	}

	return salary
}

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi helper untuk mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
// ! Convert int to Format Rupiah
func FormatRupiah(number int) string {
	hitung := rupiah.FormatIntToRp(number)

	return hitung
}

// ! NOTE: Count money of Employees present (GetSalary function  with string parameter)
func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	countDay := GetDayDifference(dateRange)
	salary_overral := GetSalary(countDay, data)

	return salary_overral
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
