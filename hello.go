package main

import "fmt"

func main() {
	/*
	   komentar kode
	   menampilkan pesan hello world
	*/
	fmt.Println("Hello world")

	// Ini adalah komentar

	fmt.Println("Hello world")

	var nama string = "Taufik"
	var angka int = 1

	fmt.Printf("Halo %s \n", nama)
	fmt.Printf("Angka %d \n", angka)

	fmt.Println("Hello ", nama)

	nama_saya := "Taufik"

	fmt.Printf("Halo %s Menggunakan   \n", nama_saya)

	var benar bool = false

	fmt.Println(benar)

	const database string = "aplikasi_cuti"

	fmt.Println(database)

	var pertama, kedua string = "Taufik", "Kresna"

	fmt.Println(pertama, kedua)
}
