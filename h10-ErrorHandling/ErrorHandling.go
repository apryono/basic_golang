package main

import (
	"errors"
	"fmt"
)

/*
	Error Handling,
*/

// membuat fungsi error handling
func bagi(a, b float32) (float32, error) {
	if b == 0 {
		// fungsi errors.New() untuk membuat fungsi custom
		var err = errors.New("Pembagi tidak boleh nol")
		return 0, err // Jika pembagi nya 0 , maka akan mengembalikan nilai hasil nya 0 , dan mengirim catatan error nya
	}

	if a == 0 {
		var err = errors.New("Pembilang tidak boleh nol")
		return 0, err
		/*
			Jika pembilang nya sama dengan 0 maka akan mengembalikan nilai
			hasil nya sama dengan 0, dan mengirim catatan errornya
		*/
	}

	/*
		Jika tidak validasi diatas dilewati atau tidak ada masalah ,
		maka akan mengembalikan nilai hasil pembagi a/b, dan mengembalikan catatan bahwa tidak ada error
	*/
	return a / b, nil // Jika Tidak ada Error
}

func hitungLuasSegitiga(alas, tinggi float32) (float32, error) {
	luas, err := bagi(alas*tinggi, 2)
	if err != nil {
		fmt.Println("Error, Tidak dapat Membagi karena : ", err)
		return 0, err
	}
	return luas, nil
}

func main() {
	hasil, err := hitungLuasSegitiga(5, 2)
	if err != nil {
		fmt.Println("Tidak Dapat di bagi karena : ", err)
	} else {
		fmt.Println(" Hasil Luas nya adalah : ", hasil)
	}

}
