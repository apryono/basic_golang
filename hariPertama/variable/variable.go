package main

import "fmt"

func main() {
	fmt.Println("INI PEMBELAJARAN VARIABLE")

	/* var di deklarasikan sebagai sebuah variable
	untuk penulisan variabel pada Go, ada 3 cara yaitu:
	1. Dituliskan tipe data nya
		=> contoh -> var nama string
	2. Tidak dituliskan tipe datanya
	untuk jenis ini, secara otomatis menentukan tipe data yang diberikan oleh programmer nya
	=> contoh -> nama :=
	3. Penulisan dengan constanta
		untuk membuat sebuah data konstanta
	 	=> contoh -> const phi = 3.14
	*/
	// jenis penulisan var ada beberapa cara
	//variable string
	var firstname string = "anjasmara"
	var lastname string
	lastname = "budiman"

	// variable integer
	var (
		angka  int = 10
		angka2 int = 20
	)

	//variable float
	var berkoma float32
	berkoma = 1.45

	//variabel kondisi
	var kondisi bool
	kondisi = true
	kondisi = false

	//variable konstanta
	const phi = 3.14

	//cara singkat pembuatan variable
	kondisi1 := true

	fmt.Println("Halo", firstname, lastname+"!")
	fmt.Println(angka)
	fmt.Println(angka2)
	fmt.Println(berkoma)
	fmt.Println(kondisi)
	fmt.Println(kondisi1)
	fmt.Println(phi)
}
