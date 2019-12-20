package main

import "fmt"

func main() {
	rekursif(5)
}

func rekursif(angka int) {
	fmt.Println(angka)
	if angka != 0 {
		rekursif(angka - 1)
	}

}
