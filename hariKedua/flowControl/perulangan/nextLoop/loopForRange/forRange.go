package main

import "fmt"

func main() {

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	/*
		Menampilkan array dengan looping - range
	*/
	// for i, fruit := range fruits {
	// 	fmt.Println(i, fruit)
	// }

	// untuk menghilangkan index nya atau tidak ingin ditampilkan

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
	/*
		Menampilkan array dengan loopingan biasa
	*/
	// for j := 0; j < len(fruits); j++ {
	// 	fmt.Println(j, fruits[j])
	// }
}
