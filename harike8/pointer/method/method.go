package main

import "fmt"

/*
	Method adalah sebuah fungsi yang menempel pada custom type (umumnya struct)
	Method memiliki akses ke field/property dari struct melalui receiver
*/

//
type Rectangle struct {
	width  int
	length int
}

// r hanya di gunakan sebagai variabel atau penanda untuk Rectangle
func (r Rectangle) getArea() int {
	return r.width * r.length
}

//
func (r *Rectangle) increaseSize() {
	r.width++
	r.length++
}

func main() {
	//variable rect yang menggunakan struct Rectangle
	rect := Rectangle{
		// hasil awal nya
		width:  10,
		length: 5,
	}
	/*
		memanggil rect untuk di cetak, dan akan memanggil fungsi getArea()
		pada fungsi getArea akan mengembalikan nilai hasil kali
	*/

	fmt.Println(rect.getArea()) // 50

	rect.increaseSize()
	fmt.Println(rect.width)  // 11
	fmt.Println(rect.length) // 6
}
