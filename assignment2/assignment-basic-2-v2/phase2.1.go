package main

import (
	"fmt"
	"GolangAssignment/rupiah"
)

func FormatRupiah(number int64) string {
    hitung := rupiah.FormatInt64ToRp(number)
	return hitung
}

func main (){
	fmt.Println(FormatRupiah(100000))
}