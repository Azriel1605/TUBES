package function

import (
	"TUBES/data"
	"fmt"
)

func SequencialSearchCrypto(A *data.TabCrypto, name string, showNotFound bool) int {
	var i int
	// Search for a cryptocurrency by name in the list
	for i = 0; i < data.NMAX; i++ {
		if A[i].Name == name {
			return i
		}
	}
	if showNotFound {
		fmt.Println("Crypto tidak ditemukan")
	}
	return -1
}

func LengtOfExistingData(A *data.TabCrypto) int {
	var i int
	for A[i].Name != "" && i < data.NMAX {
		i++
	}
	return i
}

func LengtOfExistingDataPorto(A *data.TabPorto) int {
	var i int
	for A[i].Name != "" && i < data.NMAX {
		i++
	}
	return i
}

func SequencialSearchPorto(P *data.TabPorto, name string) int {
	var i int
	// Search for a cryptocurrency by name in the list
	for i = 0; i < data.NMAX; i++ {
		if P[i].Name == name {
			return i
		}
	}
	fmt.Println("Crypto tidak ditemukan")
	return -1
}

func BinarySearchCrypto(A data.TabCrypto, name string) int {
	var low, mid, high int
	low = 0
	high = LengtOfExistingData(&A) - 1
	for low <= high {
		mid = (low + high) / 2
		if A[mid].Name == name {
			return mid
		} else if A[mid].Name < name {
			low = mid + 1
		} else if A[mid].Name > name {
			high = mid - 1
		}
	}

	return -1 // Data not found
}

func SelectionSortCrypto(A *data.TabCrypto) {
	var i, j, n, min int
	var temp data.TCrypto
	n = LengtOfExistingData(A) //Menampilkan banyakya data yang terisi

	for i = 0; i < n; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if A[j].Price < A[i].Price {
				min = j
			}
		}
		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}

}

func SelectionSortPorto(P *data.TabPorto) {
	var i, j, n, min int
	var temp data.TPorto
	n = LengtOfExistingDataPorto(P) //Menampilkan banyakya data yang terisi

	for i = 0; i < n; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if P[j].MarketPrice < P[i].MarketPrice {
				min = j
			}
		}
		temp = P[i]
		P[i] = P[min]
		P[min] = temp
	}
}

// func InsertionSortCrypto(A *data.TabCrypto) {
// }

func InsertionSortPorto(P *data.TabPorto) {
	var i, n, j int
	var temp data.TPorto

	n = LengtOfExistingDataPorto(P) //Menampilkan banyakya data yang terisi
	for i = 1; i < n; i++ {
		j = i - 1
		temp = P[i]
		for j >= 0 && P[j].MarketPrice > P[j+1].MarketPrice {
			P[j+1] = P[j]
			j--
		}
		P[j+1] = temp
	}
}