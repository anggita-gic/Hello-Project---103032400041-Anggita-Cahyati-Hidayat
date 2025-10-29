package main

import "fmt"

func menuDataBarang(tBarang *tabBarang, nBarang *int) {
	// Fungsi untuk menampilkan menu data barang
	fmt.Println("Selamat datang di Sistem Manajemen Barang!")
	fmt.Println("Silakan pilih menu berikut untuk melanjutkan.")
	var pilih int
	var x int // ini untuk no barang karena no barang itu int

	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("	   D A T A  B A R A N G   		")
		fmt.Println("-----------------------------------------")
		fmt.Println("1. TAMBAH DATA BARANG ")
		fmt.Println("2. EDIT DATA BARANG ")
		fmt.Println("3. HAPUS DATA BARANG ")
		fmt.Println("4. CETAK DATA BARANG ")
		fmt.Println("5. CARI BARANG BERDASARKAN NO BARANG ")
		fmt.Println("6. KATEGORI BARANG")
		fmt.Println("7. KEMBALI KE MENU UTAMA")
		fmt.Print("Pilih [1/2/3/4/5]: ")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahBarang(tBarang, nBarang)
		case 2:
			fmt.Println("No Barang berapa? ")
			fmt.Scan(&x)
			editBarang(tBarang, *nBarang, x)
		case 3:
			fmt.Println("No Barang berapa? ")
			fmt.Scan(&x)
			hapusBarang(tBarang, nBarang, x)
		case 4:
			cetakBarang(*tBarang, *nBarang)
		case 5:
			fmt.Println("Masukkan No Barang yang ingin dicari: ")
			var x int
			fmt.Scan(&x)
			if cari_barang_binary(*tBarang, *nBarang, x) != -1 {
				fmt.Println("Barang ditemukan.")
			} else {
				fmt.Println("Barang tidak ditemukan.")
			}
		case 6:
			seqSearch_susunKategori(*tBarang, *nBarang)
		}
		if pilih == 7 {
			break
		}
	}
}

func tambahBarang(B *tabBarang, n *int) {
	fmt.Println("Tambah barang")

	if *n <= NMAX {
		fmt.Print("NO Barang: ") // no brang sama kayak id
		fmt.Scan(&B[*n].noBarang)
		fmt.Print("Nama Barang: ")
		fmt.Scan(&B[*n].nama)
		fmt.Print("Harga Jual Barang: ")
		fmt.Scan(&B[*n].harga_jual)
		fmt.Print("Harga Beli Barang: ")
		fmt.Scan(&B[*n].harga_beli)
		fmt.Print("Stok Barang: ")
		fmt.Scan(&B[*n].stok)
		fmt.Print("Kategori Barang: ")
		fmt.Scan(&B[*n].kategori)
		*n++
	}

}

func editBarang(B *tabBarang, n int, x int) {
	var k int
	fmt.Println("Edit barang")
	k = seqSearchNoBarang(*B, n, x)
	if k != -1 {
		fmt.Scan(&B[k].noBarang, &B[k].nama, &B[k].harga_jual, &B[k].stok, &B[k].kategori, &B[k].harga_beli)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func seqSearchNoBarang(B tabBarang, n int, x int) int {
	var i int
	var idx int

	idx = -1
	i = 0
	for i < n && idx == -1 {
		if B[i].noBarang == x {
			idx = i
		}
		i++
	}
	return idx
}

func hapusBarang(B *tabBarang, n *int, x int) {
	fmt.Println("Hapus barang")
	var i, k int
	k = seqSearchNoBarang(*B, *n, x)
	if k != -1 {
		for i = k + 1; i < *n; i++ {
			B[i-1] = B[i]
		}
		*n--
		fmt.Println("Transaksi berhasil dihapus.")
	} else {
		fmt.Println("data tidak ditemukan.")
	}
}

func cetakBarang(B tabBarang, n int) {
	fmt.Println("Barang tidak ditemukan.")

	for i := 0; i < n; i++ {
		fmt.Printf("%d %s %d %d %d %s %d\n", B[i].noBarang, B[i].nama, B[i].harga_beli, B[i].harga_jual, B[i].stok, B[i].kategori, B[i].barang_terjual)
	}
}

func cari_barang_binary(T tabBarang, n int, x int) int {
	var left, right, mid int
	left = 0
	right = n - 1
	mid = (left + right) / 2

	for left <= right && T[mid].noBarang != x {
		if x < T[mid].noBarang {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	if left > right {
		return -1 // Not found
	} else {
		return mid // Found at index mid
	}
}

func seqSearch_susunKategori(B tabBarang, n int) {
	var kategoriList = [6]string{"MAKANAN", "ELEKTRONIK", "KECANTIKAN", "OTOMOTIF", "KESEHATAN", "PAKAIAN"}

	for k := 0; k < 6; k++ {
		kategori := kategoriList[k]
		adaData := false

		// Cek dulu apakah ada data untuk kategori ini
		for i := 0; i < n; i++ {
			if B[i].kategori == kategori {
				adaData = true
				break
			}
		}

		if adaData {
			fmt.Println("Kategori:", kategori)
			for i := 0; i < n; i++ {
				if B[i].kategori == kategori {
					fmt.Println("  Nama:", B[i].nama, "- Harga:", B[i].harga_jual)
				}
			}
		}
	}
}
