/*
	tanda & ini untuk menunjukkan alamat memori
*/

package main

import "fmt"

func main() {
	var num int
	var ptr *int // deklarasi ptr sebagai pointer ke integer

	num = 10
	ptr = &num // & operator digunakan untuk
	fmt.Println(*ptr)

	// menampilkan passing by value dan passing by reference
	v1 := "v1"
	v2 := "v2"

	var1 := v1
	var2 := &v2

	learnPointer(var1, var2)
	fmt.Println("var1 : ", var1)
	fmt.Println("var2 : ", var2)
	fmt.Println("var2 : ", *var2)

	// passing by reference
	nilai := 5
	angka := &nilai
	angka2 := nilai
	passByReference(angka)
	passByValue(angka2)
	// untuk menampilkan nilai nya harus di buat tanda * sebelum angka , seperti *angka, jika tidak hanya menampilkan alamat memori
	fmt.Println("angka : ", angka)
	//passing by value, menembalikan nilai yang sama
	fmt.Println("angka2 : ", angka2)

}

// membuat fungsi pointer yang memanggil dua type string biasa dan pointer string
func learnPointer(var1 string, var2 *string) {
	var1 = "var1"
	*var2 = "var2"
}

// passing by reference harus menggunakan tanda * sebagai penunjuk atau pointer
func passByReference(angka *int) {
	*angka = 10
}

// passing by value tidak mengubah data yang dikirim, hanya mengembalikan nilai yang sama
func passByValue(angka int) {
	angka = 10
}
