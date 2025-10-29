/*

// Program ini ditambahkan komentar oleh Anggita - 103032400041
menu utama
1. nambahin barang
2. ngeliat data transaksi barang (ngeliat penjualan)
3. buat beli barang

*/

package main

import "fmt"

// STRUCT BARANG
type barang struct {
	nama, kategori                                         string
	noBarang, harga_jual, stok, harga_beli, barang_terjual int
}

//no barang int supaya bisa dicari pakai binary search

//ini bebas kalian mau deklar apa aja, ini contoh aja
const NMAX int = 10

type tabBarang [NMAX]barang

var tBarang tabBarang
var nBarang int = 0

//STRUCT TRANSAKSI
type transaksi struct {
	namaBarang             string
	hargaBarang, total     float64
	noBarang, jumlahBarang int
}

//no barang int supaya bisa dicari pakai binary search

type tabTransaksi [NMAX]transaksi

var dataTransaksi tabTransaksi
var jumlahTransaksi int = 0

func main() {
	var pilih int
	var dataBarang tabBarang
	var dataTransaksi tabTransaksi
	var nBarang, jumlahTransaksi int
	fmt.Println("Selamat datang di Sistem Manajemen Barang!")
	fmt.Println("Silakan pilih menu berikut untuk melanjutkan.")
	for {
		fmt.Println("--------------------------------------------")
		fmt.Println("	  M E N U  U T A M A  	")
		fmt.Println("--------------------------------------------")
		fmt.Println("1. MENAMBAHKAN DATA BARANG ")
		fmt.Println("2. DATA TRANSAKSI PENJUALAN BARANG ")
		fmt.Println("3. KELUAR ")
		fmt.Print("Pilih [1/2/3]: ")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			menuDataBarang(&dataBarang, &nBarang) // Panggil fungsi untuk menambah data barang
		case 2:
			menuTransaksi(&dataBarang, &dataTransaksi, &nBarang, &jumlahTransaksi) // Panggil fungsi untuk transaksi
		case 3:
			fmt.Println("Terima kasih! Sampai jumpa lagi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

//harga jual
//harga beli
//stok
//terjual
//kategori
