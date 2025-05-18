package function // DONE

import (
	"TUBES/data"
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

func RmPortofolio(A *data.TabCrypto,
	P *data.TabPorto,
	name string) {
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
	fmt.Println("+------------------------------------------------------------+")
	fmt.Println("|                         PORTOFOLIO                         |")
	fmt.Println("+------------------------------------------------------------+")
	fmt.Println("|       Crypto        |   Q   |  Harga Beli  |   Harga Jual  |")
	fmt.Println("+------------------------------------------------------------+")
	for P[i].Name != "" && i < data.NMAX {
		fmt.Printf("| %-19s | %-5d | %-12.0f | %-13.0f |\n", P[i].Name, P[i].Quantity, P[i].BuyPrice, P[i].MarketPrice)
		i++
	}
	fmt.Println("")
	fmt.Println("+------------------------------------------------------------+")

}
