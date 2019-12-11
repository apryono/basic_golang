package main

import "fmt"

/*
	Materi mengenai kondisi
	if menjalankan boolean
	if condition == true {
		perintah di eksekusi (statement/command/dll)
	} else {
		kerjakan yang lain (statement/command/dll)
	}

	Apabila lebih dari satu kondisi harus menggunakan logical operator
	Logical Operator => && dan ||

*/
func main() {
	a := 5

	fmt.Println("sebelum if")
	//Materi Mengenai if, else, dan else if
	if a < 10 && a >= 5 {
		fmt.Println("Perintah di eksekusi")
	} else if a == 5 {
		fmt.Println("Perintah tidak di eksekusi")
	} else {
		fmt.Println("Final Else")
	}
	fmt.Println("Setelah if")
}
