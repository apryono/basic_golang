//komentar di golang terdiri atas dua inline dan multiline
//setiap file program harus memiliki package minimal satu yaitu package main
package main

import "fmt"

//fungsi ini yang akan pertama kali di eksekusi
func main() {
	//digunakan untuk memunculkan text ke layar(konteks ini terminal atau CMD)
	/*
		fmt.Print()-> untuk menampilkan text ke layar tanpa baris baru
		fmt.Println() -> untuk menampilkan text ke layar dengan menambahkan baris baru
		fmt.Printf()-> untuk menampilkan text ke layar dan dapat membaca formatting seperti %s, %g, dan lainnya
	*/
	fmt.Println("Hello World!")
}
