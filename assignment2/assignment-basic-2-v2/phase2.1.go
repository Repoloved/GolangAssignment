package main

import (
	"fmt"
	"GolangAssignment/rupiah"
)

func FormatRupiah(number int) string {
    hitung := rupiah.FormatIntToRp(number)
	return hitung
}

func main (){
	fmt.Println(FormatRupiah(100000))
}