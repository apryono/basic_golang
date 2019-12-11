package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		var angka1, angka2, hasil int
		fmt.Print("Masukkan Angka 1 : ")
		fmt.Scan(&angka1)
		fmt.Print("Masukkan Angka 2 : ")
		fmt.Scan(&angka2)

		hasil = angka1 + angka2
		fmt.Println("Hasil yang di masukkan adalah ", hasil)
	*/

	var kal string
	fmt.Print("Masukkan Kalimat : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kal = scanner.Text()
	fmt.Println("Kalimat yang dimasukkan adalah : ", kal)

}
