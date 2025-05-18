package function

import (
	"TUBES/data"
	"TUBES/v"
	"fmt"
)

func Beli(A *data.TabCrypto) {
	var i int
	var name string
	var quantity int
	fmt.Println("Masukkan nama crypto yang ingin dibeli: ")
	fmt.Scan(&name)
	i = SequencialSearchCrypto(A, name, false)
	if i != -1 {
		fmt.Println("Masukkan jumlah crypto yang ingin dibeli: ")
		fmt.Scan(&quantity)
		if v.Saldo >= (A[i].Price * float64(quantity)) {
			v.Saldo -= (A[i].Price * float64(quantity))
			// Masukkan transaksi ke hisotory
			HistoryRecord(&v.History, &v.IdxHistory, name, A[i].Price, quantity, "buy")
			AddPortofolio(A, &v.Porto, name, quantity)
			fmt.Println("Transaksi berhasil")
		} else {
			fmt.Println("Saldo tidak cukup")
		}
	} else {
		fmt.Println("Crypto tidak ditemukan")
	}
}

// 
func Jual(P *data.TabPorto) {
	var i int
	var name string
	var quantity int
	fmt.Println("Masukkan nama crypto yang ingin dijual: ")
	fmt.Scan(&name)
	i = SequencialSearchPorto(P, name)
	if i != -1 {
		fmt.Println("Masukkan jumlah crypto yang ingin dijual: ")
		fmt.Scan(&quantity)
		if P[i].Quantity >= quantity {
			v.Saldo += (P[i].MarketPrice * float64(quantity))
			// Masukkan transaksi ke hisotory
			HistoryRecord(&v.History, &v.IdxHistory, name, P[i].MarketPrice, quantity, "sell")
			fmt.Println("Transaksi berhasil")
		} else {
			fmt.Println("Jumlah crypto yang ingin dijual melebihi jumlah crypto yang dimiliki")
		}
	} else {
		fmt.Println("Crypto tidak ditemukan")
	}
}
