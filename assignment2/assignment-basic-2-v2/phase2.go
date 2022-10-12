package main

import (
	"fmt"
	"time"
	"math"
	"strings"
	"GolangAssignment/rupiah"
)

const TextDate = "2 January 2006"

// ! NOTE: Count Days
func GetDayDifference(date string) int {
	splitmin := strings.Split(date, " - ")
	splitspace := strings.Split(date, " ")
	start_date := splitmin[0] + " " + splitspace[5]
	end_date := splitmin[1]
	date1, _ := time.Parse(TextDate, start_date)
	date2, _ := time.Parse(TextDate, end_date)
	days := int(math.Ceil(date2.Sub(date1).Hours() / 24)+1)
	return days
}

func FormatRupiah(number int) string {
	hitung := rupiah.FormatIntToRp(number)
	
	return hitung
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
    var moneyandi, moneyRojaki, moneyraji, moneysupri, capPerson int
	var person1, person2, person3, person4 string
    for i := 0; i < rangeDay; i++ {
		capPerson = cap(data[i])
		for j := 0; j < capPerson; j++ {
			hasil := data [i][j]
            person1 = "andi"
			person2 = "Rojaki"
			person3 = "raji"
			person4 = "supri"
            if strings.Contains(hasil, person1){
				moneyandi += 50000
				} else if strings.Contains(hasil, person2){
					moneyRojaki += 50000
					} else if strings.Contains(hasil, person3){
						moneyraji += 50000
						} else if strings.Contains(hasil, person4){
							moneysupri += 50000
						}
		}
	}
	gajiandi := FormatRupiah(moneyandi)
	gajiRojaki := FormatRupiah(moneyRojaki)
	gajiraji := FormatRupiah(moneyraji)
	gajisupri := FormatRupiah(moneysupri)
    totalgaji := map[string]string{person1: gajiandi, person2: gajiRojaki, person3: gajiraji, person4: gajisupri}
    return totalgaji
}

func main() {
	res := GetSalary(3,[][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	fmt.Println(res)
}