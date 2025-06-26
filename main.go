package main

import (
	"fmt"
)

func main() {

	// cara menghitugan bunga berjalan
	var input int
	var cases1 int
	var cases2 int

	fmt.Print("input: ")
	fmt.Scan(&input)
	fmt.Print("cases1: ")
	fmt.Scan(&cases1)
	fmt.Print("cases2: ")
	fmt.Scan(&cases2)

	hasil := BungaBerjalanCalculate(input, cases1, cases2)
	fmt.Println("Bunga Berjalan:", hasil, "bulan")

}

func DibeliBiayaTunai(hasilCalculate, biayaProvisi float64) float64 {
	return hasilCalculate + biayaProvisi
}

func DiJualBiayaTunai(hasilCalculate, biayaProvisi float64) float64 {
	return hasilCalculate - biayaProvisi
}

func CalculateSaham(persenKurs, lembarSaham, priceLembar float64) float64 {
	hasil := (persenKurs / 100) * lembarSaham * priceLembar
	return hasil
}

func CalculateObligation(persenKurs, lembarObligation, pricelembar float64) float64 {
	hasil := (persenKurs / 100) * lembarObligation * pricelembar
	return hasil
}

func BungaBerjalanCalculate(input, cases1, cases2 int) int {
	bulan := 12
	hasil := 0
	if input < cases1 && input < cases2 {
		hasil = bulan - cases2 + input
	} else if input == cases1 && input == cases2 {
		hasil = 0
	} else if input > cases1 {
		hasil = input - cases1
	}

	return hasil
}

// Pilih Menu
// 1. Saham
// a. dijual
// b. dibeli
// 2. Obligation
// a. dijual
// b. dibeli
// 3. Exit

// mencari harga netto
