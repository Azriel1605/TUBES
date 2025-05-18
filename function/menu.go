package function

import "fmt"

func MainMenu() {
	fmt.Println("+------------------------------+")
	fmt.Println("|             MENU             |")
	fmt.Println("+------------------------------+")
	fmt.Println("| 1. Modifikasi Aset           |") //Done
	fmt.Println("| 2. Transaksi                 |") //Done
	fmt.Println("| 3. Portofolio                |") //Done
	fmt.Println("| 4. Histori Pembelian         |") //Done
	fmt.Println("| 5. Cari Aset Crypto          |") //Done
	fmt.Println("| 6. Urutkan Aset (harga)      |") //Done
	fmt.Println("| 7. Urutkan Portofolio (harga)|") //Done
	fmt.Println("| 8. Tampilkan Semua Aset      |") //Done
	fmt.Println("| 9. Exit                      |")
	fmt.Println("+------------------------------+")
}

func ModifyCryptoMenu() {
	fmt.Println("\n+-----------------------------+")
	fmt.Println("|-_-_- Modifikasi Crypto -_-_-|")
	fmt.Println("+-----------------------------+")
	fmt.Println("| 1. Tambah Crypto            |")
	fmt.Println("| 2. Hapus Crypto             |")
	fmt.Println("| 3. Ubah Harga Crypto        |")
	fmt.Println("| 4. Kembali ke Menu Utama    |")
	fmt.Println("+-----------------------------+")
}

func TransaksiMenu() {
	
	fmt.Println("\n+----------------------------+")
	fmt.Println("| 1. Beli Crypto             |")
	fmt.Println("| 2. Jual Crypto             |")
	fmt.Println("| 3. Kembali ke Menu Utama   |")
	fmt.Println("+----------------------------+")

}

