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

func GetSalaryOverview(dateRange string, data [][]string)map[string]string {
	countDay := GetDayDifference(dateRange)
    var moneyandi int
	var moneyRojaki int
	var moneyraji int
	var moneysupri int
    for i := 0; i < countDay; i++ {

			hasil := data [i][j]
            person1 := "andi"
			person2 := "Rojaki"
			person3 := "raji"
			person4 := "supri"
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
    uangandi := FormatRupiah(moneyandi)
	uangRojaki := FormatRupiah(moneyRojaki)
	uangraji := FormatRupiah(moneyraji)
	uangsupri := FormatRupiah(moneysupri)
    gaji := map[string]string{"andi": uangandi, "Rojaki": uangRojaki, "raji": uangraji, "supri": uangsupri}
    return gaji
}

func main() {
	res := GetSalaryOverview("21 February - 23 February 2021",[][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	fmt.Println(res)
}