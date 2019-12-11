/*
	Ini Merupakan Materi Operator Artimetika
*/
package main

import "fmt"

func main() {
	a := 10
	b := 10 + 5
	c := b - a
	d := b * c
	e := 7 % 3

	fmt.Println("Nilai a =", a)
	fmt.Println("Nilai b =", b)
	fmt.Println("Nilai c =", c)
	fmt.Println("Nilai d =", d)
	fmt.Println("Nilai e =", e)
	//e = e + 1 atau e++ untuk menambahkan satu nilai
	e += 1 // ini sama dengan e = e + 1
	fmt.Println("Nilai e =", e)
}
