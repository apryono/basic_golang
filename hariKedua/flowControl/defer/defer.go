package main

import "fmt"

func main() {

	// defer digunakan untuk menunda program, dan akan di tampilkan setelah program lain di jalan kan
	defer fmt.Println("World")

	fmt.Print("Hello ")
}
