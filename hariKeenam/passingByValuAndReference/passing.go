package main

import "fmt"

func main() {
	var a = 10
	tambahDua(a)
	fmt.Println(a)
}

func tambahDua(angka int) {
	angka = angka + 2
}
