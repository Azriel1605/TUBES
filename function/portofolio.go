package function // DONE

import (
	"TUBES/data"
	"TUBES/v"
	"fmt"
)

func AddPortofolio(A *data.TabCrypto,
	P *data.TabPorto,
	name string,
	quantity int) {
	var i, j, idxC int
	i = SequencialSearchPorto(P, name)
	if i != -1 {
		P[i].Quantity += quantity

	} else {
		idxC = SequencialSearchCrypto(A, name, false)
		for P[j].Name != "" && j < data.NMAX {
			j++
		}
		P[j].Name = name
		P[j].Quantity = quantity
		P[j].BuyPrice = A[idxC].Price
		P[j].MarketPrice = A[idxC].Price
	}
}

func SellPortofolio(P *data.TabPorto, name string, n int){
	// Ngurangin Lot di portofolio
	// 1. Cari index di portofolio -> kalau ga ada print(tidak ditemukan)
	// 2. Kalau ada, jika n > P[i].quantity, berati lot nya lebih dari yang ada di portofolio
	// 3. Kalau ga minus, Kurangin stock lot nya
	// 4. Kalau udah 0, hapus dari portofolio

	// Code nya ditulis disini
	// bentar bang mak gw neplon
	// sok ae, gw sholat dl

	var i int = SequencialSearchPorto(P, name)

	if i == -1 { // Kalau tidak ada nama crypto di portofolio
		fmt.Println("Tidak ditemukan")
	} 
		
	if i != -1 { // Kalau ada nama di portofolio
		if P[i].Quantity < n {
				fmt.Println("Lot lebih Banyak dari yang ada di portofolio")
		} else {
			P[i].Quantity -= n
			if P[i].Quantity == 0 {
				P[i].Name = ""
				P[i].Quantity = 0
				P[i].BuyPrice = 0
				P[i].MarketPrice = 0
			}
			fmt.Println("Crypto dalam portofolio berhasil dijual")
		}
	}
}

func RmPortofolio(P *data.TabPorto,
	name string) {

	// Ini buat ngapusnya
	var i int = SequencialSearchPorto(P, name)
	if i != -1 {
		P[i].Name = ""
		P[i].Quantity = 0
		P[i].BuyPrice = 0
		P[i].MarketPrice = 0
	}
}

func ShowPortofolio(P *data.TabPorto) {
	//Looping index  NMAX
	//fmtprint P[i].name, .quantiti etc etc
	var i int
	var equity int

	fmt.Println("+------------------------------------------------------------+")
	fmt.Println("|                         PORTOFOLIO                         |")
	fmt.Println("+------------------------------------------------------------+")
	fmt.Println("|       Crypto        |   Q   |  Harga Beli  |   Harga Jual  |")
	fmt.Println("+------------------------------------------------------------+")
	for P[i].Name != "" && i < data.NMAX {
		fmt.Printf("| %-19s | %-5d | %-12.0f | %-13.0f |\n", P[i].Name, P[i].Quantity, P[i].BuyPrice, P[i].MarketPrice)
		equity += P[i].Quantity * int(P[i].MarketPrice)
		i++
	}
	equity += int(v.Saldo)
	fmt.Println("+------------------------------------------------------------+")
	fmt.Printf("| %-19s : %36.0f |\n", "Sisa Modal", float64(v.Saldo))
	fmt.Println("+------------------------------------------------------------+")
	fmt.Printf("| %-19s : %36.0f |\n", "Total Ekuitas", float64(equity))
	fmt.Println("+------------------------------------------------------------+")

}
