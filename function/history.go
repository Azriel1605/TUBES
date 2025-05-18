package function

import (
	"TUBES/data"
	"fmt"
)

func HistoryRecord(H *data.TabHistory, idx *int, name string, price float64, quantity int, transactionType string) {
	H[*idx].Name = name
	H[*idx].BuyPrice = price
	H[*idx].Quantity = quantity
	H[*idx].TotalPrice = price * float64(quantity)
	H[*idx].TransactionType = transactionType
	*idx++
}

func ShowHistory(H *data.TabHistory) {
	var i int
	fmt.Println("+---------------------------------------------------------------------------+")
	fmt.Println("|                             History Transaksi                             |")
	fmt.Println("+---------------------------------------------------------------------------+")
	fmt.Println("|Name           |Buy Price    |Quantity |Total Transaction |Transaction Type|")
	for H[i].Name != "" && i < data.NMAX {
		fmt.Printf("|%-14s |%-12.0f |%-8d |%-17.0f |%-15s |\n",
			H[i].Name, H[i].BuyPrice, H[i].Quantity, H[i].TotalPrice, H[i].TransactionType)
		i++
	}
	fmt.Println("+---------------------------------------------------------------------------+")
}
