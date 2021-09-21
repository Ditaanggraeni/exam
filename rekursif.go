package main

import (
	"fmt"
	"math"
	"strings"
)

var total = 0
var sum = 0

func deretAngka(jumlah int, awal int) int {

	if jumlah <= 1 {
		fmt.Printf("%d + ", awal)
		total = total + awal
		return awal
	}

	i := deretAngka(jumlah-1, awal)
	pangkat := int(math.Pow(float64(10), float64(jumlah-1)))
	total = awal * pangkat
	fmt.Printf("%d + ", total)

	return i + total
}

func deretPecahan(i int, jumlah int, pecahan string) int {

	if i <= jumlah {
		pangkat := int(math.Pow(float64(i), float64(2)))
		if pecahan == "penyebut" {
			fmt.Printf(" x^%d/%d + ", i, pangkat)
		} else {
			fmt.Printf("%dx/%d  ", pangkat, i)
		}

		return pangkat + deretPecahan(i+1, jumlah, pecahan)

	} else {
		return 0
	}
}

func upperCase(kata string, i int) int {
	upper := strings.ToUpper(kata)

	if i < len(kata) {
		var filter bool
		if string(kata[i]) == string(upper[i]) {
			filter = true
		}
		if filter == true {
			return 1 + upperCase(kata, i+1)
		} else {
			return 0 + upperCase(kata, i+1)
		}
	} else {
		return 0
	}
}

func main() {
	var jumlah int
	var kata string

	fmt.Print("Masukan jumlah deret : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukan Kata : ")
	fmt.Scan(&kata)

	fmt.Print("1. ")
	fmt.Println(" =", deretAngka(jumlah, 5))
	fmt.Print("2. ")
	fmt.Println("= x^ ", "/", deretPecahan(1, jumlah, "penyebut"))
	fmt.Print("3. ")
	fmt.Printf("Jumlah huruf kapital pada kata %s : ", kata)
	fmt.Println(upperCase(kata, 0))
}
