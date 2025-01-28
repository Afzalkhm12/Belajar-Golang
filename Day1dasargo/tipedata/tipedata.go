package tipedata

import "fmt"

func Tipedata () {
	var positiveNumber uint64 = 109
var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
}

func Desimal () {
	var decimalNumber = 3.90
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	
}

func bol () {
	var bol bool = true
	fmt.Printf("exist? %t \n", bol)
}

func str () {
	var str string = `Hawooo Perkenalkan my Name "Afzal".
	From Universitas Nasional, Jakarta Selatan`
	fmt.Println(str)
}