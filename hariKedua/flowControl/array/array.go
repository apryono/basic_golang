package main

import "fmt"

func main() {

	var names [5]string
	names[0] = "BUDI"
	names[1] = "BUDA"
	names[4] = "BODA"

	var fruits = [3]string{"nanas", "semangka", "anggur"}
	// apabila ingin di custom atau mengubah isi di fruits, bisa di tambah seperti di bawah
	fruits[1] = "mangga"
	fruits[2] = "jeruk"

	/*
		jika belum tau berapa isi array yang akan ditentukan ,
		bisa menggunakan tanda [...]typeData{array1,array2,array3,...} sebagai berikut:
	*/
	var numb = [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	//len fungsi nya untuk menghitung jumlah pada array
	fmt.Println(numb, len(numb))
	fmt.Println(names, len(names))
	fmt.Println(fruits)
}
