/*
	Array adalah baris atau kumpulan data
	index itu adalah nomor ,dan dimulai dari hitungan 0


*/
package main

import "fmt"

func main() {

	var arrAngka [10]int
	arrAngka[0] = 50
	arrAngka[5] = 150

	fmt.Println(arrAngka)

	//cara deklarasi array tak terbatas, dan akan menyesuaikan isi dari data yang di inisialisasikan oleh array nya
	var arrKata = [...]string{
		"Joe", "ryo", "andi", "julia",
	}
	fmt.Println(arrKata)
	//len itu untuk menghitung isi datanya, bisa untuk menghitung isi array bisa juga untuk menghitung jumlah karakter
	fmt.Println(len(arrKata))

	fmt.Print("--------------------------------------------\n")

	var arrNumber [10]int

	for i := 0; i < len(arrNumber); i++ {
		// untuk mencertak hasil perulangan
		arrNumber[i] = i
	}
	fmt.Println(arrNumber)

	fmt.Print("--------------------------------------------\n")
	// for range nya hanya menggunakan satu variabel
	for i := range arrAngka {
		fmt.Printf("Index Ke %v Nilai nya adalah %v \n", i, arrAngka[i])
	}
	fmt.Print("--------------------------------------------\n")

	// for range nya hanya menggunakan dua variabel
	for i, val := range arrNumber {
		fmt.Printf("Index Ke %v Nilai nya adalah %v \n", i, val)
	}

	fmt.Print("--------------------------------------------\n")
	// for range tanpa menampilkan nilai index nya, yaitu dengan membuat indeks i nya dengan tanda underscore "_"
	for _, val := range arrNumber {
		fmt.Printf("Nilai nya adalah %v \n", val)
	}

}
