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
	// a := 5

	// fmt.Println("sebelum if")
	// //Materi Mengenai if, else, dan else if
	// if a < 10 && a >= 5 {
	// 	fmt.Println("Perintah di eksekusi")
	// } else if a == 5 {
	// 	fmt.Println("Perintah tidak di eksekusi")
	// } else {
	// 	fmt.Println("Final Else")
	// }
	// fmt.Println("Setelah if")

	/*
		%.1f = hanya menentukan satu digit di belakang koma contoh 0.2
		%.2f = hanya menentukan satu digit di belakang koma contoh 0.20
		%f => untuk nilai yang hasil nya float
		%s => untuk nilai hasil nya string
		%v => default ,dapat di gunakan untuk keseluruhan
	*/
	var point = 6840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1v%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

}
