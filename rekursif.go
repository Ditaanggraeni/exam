package main

import (
	"fmt"
	"math"
	"strings"
)

var total = 0
var sum = 0
var jmlh = 0.0
var jumlahI = 0.0

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

func deretPecahan(i float64) float64 {
	if i <= 1 {
		jumlahI = jumlahI + 1
		jmlh = jmlh + 2
		fmt.Printf("x^1/1 + ")
		return i
	}
	y := deretPecahan(i - 1)
	jumlahI = jumlahI + i
	x := i * i
	jmlh = jmlh + x
	fmt.Printf("x^%.0f/%.0f + ", i, x)
	return y
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
	var jumlah, pilihan int
	var kata string
	var x = true

	for x {
		fmt.Println("\n================== PILIHAN ===================")
		fmt.Println("1. Deret Perkalian Angka")
		fmt.Println("2. Deret Bilangan Pecahan")
		fmt.Println("3. Hitung Jumlah Huruf Kapital Pada Kalimat")
		fmt.Println("0. Keluar")
		fmt.Println("==============================================")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scan(&pilihan)
		fmt.Println("==============================================")

		switch pilihan {
		case 1:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println(" =", deretAngka(jumlah, 5))
		case 2:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			deretPecahan(float64(jumlah))
			fmt.Printf("= x^%.0f/%.0f\n", jumlahI, jmlh)
		case 3:
			fmt.Print("Masukan Kata : ")
			fmt.Scan(&kata)
			fmt.Printf("Jumlah huruf kapital pada kata %s : ", kata)
			fmt.Println(upperCase(kata, 0))
		case 0:
			x = false
		}
	}
}
