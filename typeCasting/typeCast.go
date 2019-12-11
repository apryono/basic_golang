package main

import "fmt"

func main() {
	var i int = 97
	var s string

	fmt.Println(i, "ini juga di print")
	/*
		Sprintln digunakan untuk force int atau type data lain menjadi string
		variabel s digunakan untuk menyimpan hasil Sprintln
	*/

	s = fmt.Sprintln(i, "ini juga di print")

	fmt.Println("isi variabel i adalah ", s)
	//dapat juga ditampilkan seperti di bawah ini
	// fmt.Println("isi variabel i adalah ", fmt.Sprintln(i, "ini juga di print"))
	fmt.Println("selesai")

	/*
		Printf untuk menampilkan format yang ada di golang
		default nya yaitu %v
	*/
	a, b := 100, 200
	fmt.Printf("printf: ini %v dan ini %v\n", a, b)
	fmt.Printf("printf: ini %[2]v dan ini %[1]v", a, b)

}
