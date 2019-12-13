package main

import "fmt"

func main() {
	var i int

	/*
		Penggunaan Keyword break & continue
		Keyword break digunakan untuk menghentikan secara paksa sebuah perulangan,
		sedangkan continue dipakai untuk memaksa maju ke perulangan berikutnya.
	*/
	// for i = 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("Angka ke : ", i)
	// }

	/*
		PERULAGAN BERSARANG
		Tak hanya seleksi kondisi yang bisa bersarang, perulangan juga bisa.
		Cara pengaplikasiannya kurang lebih sama,
		tinggal tulis blok statement perulangan didalam perulangan.
	*/

	for i = 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

}
