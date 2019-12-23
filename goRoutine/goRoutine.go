package main

import (
	"fmt"
	"time"
)

//concurency, berjalan secara pararel
// fungsi chan untuk menunggu dan menerima nilai, channel tidak harus menerima nilai bool atau bisa apa aja
func jumlah(a, b int, c chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println(a, b, a+b)
	c <- true // artinya akan menerima nilai
}
func main() {
	// fmt.Println("Hello")
	// time.Sleep(1000 * time.Millisecond)
	// fmt.Println("World")

	c := make(chan bool) // untuk membuat chan
	go jumlah(5, 5, c)
	// time.Sleep(1 * time.Second)
	selesai := <-c // artinya menunggu nilai dan akan sleep sampai nilai nya ke isi di suatu tempat
	fmt.Println(selesai)

}
