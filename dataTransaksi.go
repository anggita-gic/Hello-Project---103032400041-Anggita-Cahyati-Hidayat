package main

import "fmt"

func menuTransaksi(databarang *tabBarang, dataTransaksi *tabTransaksi, nBarang *int, jumlahTransaksi *int) {
	// Fungsi untuk menampilkan menu transaksi
	fmt.Println("Selamat datang di Sistem Manajemen Transaksi!")
	fmt.Println("Silakan pilih menu berikut untuk melanjutkan.")
	var pilih int
	var x int //ini untuk no barang karena no barang itu int

	for {
		fmt.Println("--------------------------------------------")
		fmt.Println("	  D A T A  T R A N S A K S I  	")
		fmt.Println("--------------------------------------------")
		fmt.Println("1. TAMBAH DATA TRANSAKSI PENJUALAN BARANG ")
		fmt.Println("2. EDIT DATA TRANSAKSI PENJUALAN BARANG ")
		fmt.Println("3. HAPUS DATA TRANSAKSI PENJUALAN BARANG ")
		fmt.Println("4. CETAK DATA TRANSAKSI PENJUALAN BARANG ")
		fmt.Println("5. DATA MODAL DAN PENDAPATAN ")
		fmt.Println("6. 5 BARANG TERLARIS PENJUALAN ")
		fmt.Println("7. KEMBALI KE MENU UTAMA")
		fmt.Print("Pilih [1/2/3/4/5]: ")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahTransaksi(dataTransaksi, databarang, jumlahTransaksi, nBarang)
		case 2:
			fmt.Println("No Barang berapa? ")
			fmt.Scan(&x)
			editTransaksi(dataTransaksi, *jumlahTransaksi, x)
		case 3:
			fmt.Println("No Barang berapa? ")
			fmt.Scan(&x)
			hapusTransaksi(dataTransaksi, jumlahTransaksi, x)
		case 4:
			cetakTransaksi(*dataTransaksi, *jumlahTransaksi)
		case 5:
			hitungTotalHarga(*dataTransaksi, databarang, *jumlahTransaksi)
		case 6:
			barang_terlaris_insertionsort(databarang, *nBarang) // Menampilkan 5 barang terlaris
		}
		if pilih == 7 {
			break
		}
	}
}

// untuk menambahkan data transaksi
func tambahTransaksi(T *tabTransaksi, B *tabBarang, n, nBarang *int) {
	if *n < NMAX {
		fmt.Println("Tambah Transaksi")
		fmt.Print("No Barang: ")
		fmt.Scan(&T[*n].noBarang)
		//ini di check dlu apakah data barang sudah ada atau belum
		if cari_barang_binary(*B, *nBarang, T[*n].noBarang) == -1 {
			fmt.Println("Barang tidak ditemukan. Silakan tambahkan barang terlebih dahulu.")
		} else {
			fmt.Print("Jumlah Barang: ")
			fmt.Scan(&T[*n].jumlahBarang)
			T[*n].total = T[*n].hargaBarang * float64(T[*n].jumlahBarang)
			*n++
			B[*n].barang_terjual += T[*n].jumlahBarang // Update jumlah barang terjual
		}
	} else {
		fmt.Println("Data transaksi penuh")
	}
}

// fungsi ini untuk mengedit Data Transaksi
func editTransaksi(T *tabTransaksi, n int, x int) {
	var k int
	fmt.Println("Edit Transaksi")

	k = seqSearchNoBarang_transaksi(*T, n, x)
	if k != -1 {
		fmt.Print("No Barang Baru: ")
		fmt.Scan(&T[k].noBarang)
		fmt.Print("Nama Barang Baru: ")
		fmt.Scan(&T[k].namaBarang)
		fmt.Print("Harga Barang Baru: ")
		fmt.Scan(&T[k].hargaBarang)
		fmt.Print("Jumlah Barang Baru: ")
		fmt.Scan(&T[k].jumlahBarang)

		T[k].total = T[k].hargaBarang * float64(T[k].jumlahBarang)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

// untuk mencari data barang transaksi berdasarkan nomor barang
// dengan Sequential Search
func seqSearchNoBarang_transaksi(T tabTransaksi, n int, x int) int {
	var i int
	var idx int

	idx = -1
	i = 0
	for i < n && idx == -1 {
		if T[i].noBarang == x {
			idx = i
		}
		i++
	}
	return idx
}

// fungsi hapus Data Transaksi
func hapusTransaksi(T *tabTransaksi, n *int, x int) {
	fmt.Println("Hapus transaksi")
	var i, k int
	k = seqSearchNoBarang_transaksi(*T, *n, x)
	if k != -1 {
		for i = k + 1; i < *n; i++ {
			T[i] = T[i+1]
		}
		*n--
		fmt.Println("Transaksi berhasil dihapus.")
	} else {
		fmt.Println("data tidak ditemukan.")
	}

}

// untuk menghitung total harga dari transaksi
func hitungTotalHarga(T tabTransaksi, B *tabBarang, n int) {
	var total int
	var i int

	//mencari data modal suatu  barang
	total = 0
	for i = 0; i < n; i++ {
		total += B[i].harga_beli * (B[i].stok)
	}
	fmt.Printf("Total Modal: %.2f\n", float64(total))

	//mencari data pendapatan suatu barang
	var pendapatan int
	pendapatan = 0
	for i = 0; i < n; i++ {
		pendapatan += (B[i].harga_jual) * (T[i].jumlahBarang)
	}
	fmt.Printf("Total Pendapatan: %.2f\n", float64(pendapatan))
	fmt.Printf("Keuntungan: %.2f\n", float64(pendapatan-total))
}

// untuk mencetak data transaksi
func cetakTransaksi(T tabTransaksi, n int) {
	fmt.Println("Cetak barang")

	for i := 0; i < n; i++ {
		fmt.Printf("%d %s %.2f %d\n", T[i].noBarang, T[i].namaBarang, T[i].hargaBarang, T[i].jumlahBarang)
	}
}

// fungsi ini untuk mencari barang terlaris berdasarkan banyaknya jumlah barang
func barang_terlaris_insertionsort(tabBarang *tabBarang, n int) {
	for i := 1; i < n; i++ {
		key := (*tabBarang)[i]
		j := i - 1
		for j >= 0 && (*tabBarang)[j].barang_terjual < key.barang_terjual {
			(*tabBarang)[j+1] = (*tabBarang)[j]
			j--
		}
		(*tabBarang)[j+1] = key
	}

	fmt.Println("5 Barang Terlaris:")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Printf("%d - %s - Terjual: %d\n", (*tabBarang)[i].noBarang, (*tabBarang)[i].nama, (*tabBarang)[i].barang_terjual)
	}

}
