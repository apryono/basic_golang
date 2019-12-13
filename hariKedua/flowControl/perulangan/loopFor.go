package main

/*
	Di dalam perulangan for yang pertama kali dilakukan yaitu mengecek dari kiri ke kanan
	for statement1; statement2; statement3 {
		statement4
	}
	contoh : for i := 1; i < 10; i++ {
					fmt.Println(i)
				}
				yang pertama kali di cek yaitu i := 1
				kemudian lanjut ke i < 10, apabila hasil nya true
				kemudian akan di tampilkan dahulu ke fmt.Println(i).
				setelah itu lanjut ke i++. Kemudian di cek kembali ke statement2
				Secara berulang sampai memenuhi statement2
*/
import "fmt"

func main() {
	/* --------- Perulangan For -------- */
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	/* --------- Perulangan For a While --------
		For yang konsep nya mirip seperti While
	   -----------------------------------------
	*/
	// i := 0

	// for i < 5 {
	// 	fmt.Println("Success", i)
	// 	i++
	// }

	/* --------- Perulangan For Tanpa Argumen --------
		Cara ke-3 adalah for ditulis tanpa kondisi.
		Dengan ini akan dihasilkan perulangan tanpa henti (sama dengan for true).
		Pemberhentian perulangan dilakukan dengan menggunakan keyword break
	   -----------------------------------------
	*/

	i := 0
	for {
		fmt.Println("Angka", i)

		i++
		if i == 5 {
			break
		}
	}

}
