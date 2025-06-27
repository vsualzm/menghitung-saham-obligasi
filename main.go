package main

func main() {

	// cara menghitugan bunga berjalan
	// var input int
	// var cases1 int
	// var cases2 int
	// fmt.Print("input: ")
	// fmt.Scan(&input)
	// fmt.Print("cases1: ")
	// fmt.Scan(&cases1)
	// fmt.Print("cases2: ")
	// fmt.Scan(&cases2)
	// hasil := BungaBerjalanCalculate(input, cases1, cases2)
	// fmt.Println("Bunga Berjalan:", hasil, "bulan")

}

// penyelesaian soal saham
// 1. di beli
// harga = harga kurs * jumlah lembar saham yg di beli * harga perlembar saham
// harga di bayar tunai = harga + provisi
// pembagian harga saham perlembar
// Jurnal nya
// Dr. Marketable securities: hargadibayartunai
// Cr.             Cash     :					hargadibayartunai

// 2. dijual
// harga = harga kurs * jumlah lembar saham yg di jual * harga perlembar saham
// harga di bayar tunai = harga - provisi
// harga pokok penjualan saham = jumlah lembar saham yg di jual * harga perlembar saham
// mencari laba = harga di bayar tunai - harga pokok penjualan saham
// Jurnal nya
// Dr. Cash                 				 : diterima tunai
// Cr.             M/S      				 :               hargaPokokPenjualanSaham
// Cr.             Profit On Sales securities:				 laba

// penyelesaian soal obligation
// 1. di beli
// harga = harga kurs * jumlah lembar saham yg di jual * harga perlembar saham
// nilaiMS = harga + provisi
// Bunga berjalan
// if bulanBeli < bulan1 && input < bulan2 {
// 		hasil = bulan12 - bulan2 + bulanBeli
// 	} else if bulanBeli == bulan1 && input == bulan2 {
// 		hasil = 0
// 	} else if bulanBeli > bulan1 {
// 		hasil = bulanBeli - bulan1
// 	}

// harga2 = bulanBerjalan/12 * bungaobligasi% * jumlah lembar obligasi yg di beli * harga perlembar obligasi
// dibayartunai = harga + harga2

// mencari harga perlembar obligasi
// perlembar = nilai M/S / jumlah lembar obligasi yg di beli
// JURNAL
// Dr. Marketable securities: nilaiMS
// Dr. Interest earned      : harga2
// Cr.             Cash     :					dibayartunai

// mencari diterima bunga obligasi

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
