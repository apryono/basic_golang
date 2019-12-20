package main

func main() {
	//defer akan di jalankan terlebih dahulu
	
	defer println()
	defer print("A")
	print("B")
	defer print("C")
	print("D")
}
