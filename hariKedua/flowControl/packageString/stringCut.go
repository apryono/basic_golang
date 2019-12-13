package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = ("   Jhon wick ada")
	var name = ("  catalogue")

	var numb int = 1
	result := printNumber(numb)

	cutString(data, name)

	/*
		Fungsi TrimSpace untuk memotong space di awal
	*/
	fmt.Println(result)

}

func printNumber(angka int) int {
	if angka == 1 {
		return 1
	} else {
		return 2
	}
}

func cutString(param1, param2 string) {
	fmt.Println(strings.TrimSpace(param1))
	fmt.Println(strings.TrimSpace(param2))
}
