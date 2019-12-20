package main
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

)

const DATABASE = "datasiswa.db"
​
type mahasiswa struct {

	Id      int    `json:"Id"`
	Nama    string `json:"Nama"`
	Umur    string `json:"Umur"`
	Jurusan string `json:"Jurusan"`
}

var mHs []mahasiswa

func main() {

	if fileExists(DATABASE) {

		getData := readFile(DATABASE)

		if getData != "" {

			mHs = jsonDecode(getData)

		}

	}

	var input int

	for {

		// fmt.Println("======== [ degbug ] ========")

		// fmt.Println(mHs)

		// fmt.Println("======== [ degbug ] ========")

		clearScreen()

		fmt.Println("--===================[ Main Menu ]===================--")

		fmt.Print("1. Tambah Mahasiswa\n2. Hapus Mahasiswa\n3. Lihat Mahasiswa\n4. Exit\nSilakan masukan menu yang di pilih : ")

		input, _ = strconv.Atoi(inputData())

		if input > 0 && input <= 4 {

			switch input {

			case 1:

				clearScreen()

				inputMahasiswa()

			case 2:

				clearScreen()

				deleteMahasiswa()

			case 3:

				//clearScreen()

				viewMahasiswa()

			case 4:

				fmt.Println("--===================[ Exit ]===================--")

				os.Exit(0)

			}

		} else {

			fmt.Println("--===================[ Main Menu ]===================--")

			fmt.Println("-[+] Pilih menu yang tersedia!")

			fmt.Println("--===================[ Main Menu ]===================--")

		}

​

		sleep()

		// break

	}

}

​

func inputMahasiswa() {

	if len(mHs) < 5 {

		for {

			var nama string

			var umur int

			var jurusan string

			var err error

			//nama = "carmudi"

			fmt.Println("========== [ Tambah Mahasiswa ] ==========")

			fmt.Printf("Nama (3 - 20 karakter): ")

			nama = inputData()

			typeDataNama := typeOfVar(nama)

			if typeDataNama == "string" && typeDataNama != "unknow" {

				if len(nama) >= 3 && len(nama) <= 20 {

					for {

						println("nama: " + nama)

						fmt.Printf("Umur (min 17 Tahun): ")

						umur, err = strconv.Atoi(inputData())

						if err == nil {

							if umur >= 17 {

								for {

									fmt.Println("umur: " + strconv.Itoa(umur))

									fmt.Printf("Jurusan (3 - 20 karakter): ")

									jurusan = inputData()

									if len(jurusan) > 0 && len(jurusan) <= 10 {

										id := len(mHs) + 1

										mHs = append(mHs, mahasiswa{Id: id, Nama: nama, Umur: strconv.Itoa(umur), Jurusan: jurusan})

										// pushData := []mahasiswa{{Id: id, Nama: nama, Umur: strconv.Itoa(umur), Jurusan: jurusan}}

										if createFiles(DATABASE, jsonEncode(mHs)) {

											fmt.Println("========== [ Tambah Mahasiswa ] ==========")

											fmt.Println("-[+] Success : Mahasiswa baru berhasil di tambahkan!")

											fmt.Println("========== [ Tambah Mahasiswa ] ==========")

										} else {

											fmt.Println("========== [ Tambah Mahasiswa ] ==========")

											fmt.Println("-[+] Failed : Mahasiswa baru berhasil di tambahkan!")

											fmt.Println("========== [ Tambah Mahasiswa ] ==========")

										}

										break

									} else {

										fmt.Println("========== [ Tambah Mahasiswa ] ==========")

										fmt.Println("-[+] Error : jurusan harus diisi dan kurang dari 10 karakter")

										fmt.Println("========== [ Tambah Mahasiswa ] ==========")

									}

									sleep()

									clearScreen()

								}

								break

							} else {

								fmt.Println("========== [ Tambah Mahasiswa ] ==========")

								fmt.Println("-[+] Error : umur minimal harus 17 tahun keatas")

								fmt.Println("========== [ Tambah Mahasiswa ] ==========")

							}

						} else {

							fmt.Println("========== [ Tambah Mahasiswa ] ==========")

							fmt.Println("-[+] Error : input harus berupa angka angka!")

							fmt.Println("========== [ Tambah Mahasiswa ] ==========")

						}

						sleep()

						clearScreen()

					}

					break

				} else {

					fmt.Println("========== [ Tambah Mahasiswa ] ==========")

					fmt.Println("-[+] Error : nama kurang dari 3 atau lebih dari 20 karakter")

					fmt.Println("========== [ Tambah Mahasiswa ] ==========")

				}

			} else {

				fmt.Println("========== [ Tambah Mahasiswa ] ==========")

				fmt.Println("-[+] Error : input nama harus alfabet/huruf!")

				fmt.Println("========== [ Tambah Mahasiswa ] ==========")

			}

			sleep()

			clearScreen()

		}

	} else {

		fmt.Println("========== [ Tambah Mahasiswa ] ==========")

		fmt.Println("-[+] Warning : penyimpanan sudah data penuh !")

		fmt.Println("========== [ Tambah Mahasiswa ] ==========")

	}

}

​

func deleteMahasiswa() {

	var id int

	var tmp []mahasiswa

	var dumpID string

	var err error

​

	for i := range mHs {

		dumpID += strconv.Itoa(i + 1)

	}

​

	if len(mHs) > 0 {

		for {

			clearScreen()

			fmt.Println("========== [ Hapus Mahasiswa ] ==========")

			fmt.Println("1. Hapus data by Index\n2. Hapus data terakhir\n3. Back")

			fmt.Print("Masukan Pilihan: ")

			id, err = strconv.Atoi(inputData())

			if err == nil {

				switch id {

				case 1:

					clearScreen()

					if len(mHs) > 0 {

						fmt.Println("========== [ Hapus Mahasiswa ] ==========")

						fmt.Println("List ID : ", strings.Split(dumpID, ""))

						fmt.Printf("-[+] Masukan id: ")

						id, err = strconv.Atoi(inputData())

						if err == nil {

							if id > 0 {

								if id <= len(mHs) {

									for i := 0; i < len(mHs); i++ {

										if i != id-1 {

											tmp = append(tmp, mHs[i])

										}

									}

									mHs = tmp

​

									if createFiles(DATABASE, jsonEncode(mHs)) {

										fmt.Println("========== [ Hapus Mahasiswa ] ==========")

										fmt.Println("-[+] Success : hapus berhasil !")

										fmt.Println("========== [ Hapus Mahasiswa ] ==========")

										if tryAgain("Apakah anda ingin kembali kemenu utama? (Y/n) : ") {

											break

										} else {

											continue

										}

									} else {

										fmt.Println("========== [ Hapus Mahasiswa ] ==========")

										fmt.Println("-[+] Error : hapus gagal!")

										fmt.Println("========== [ Hapus Mahasiswa ] ==========")

									}

									break

								} else {

									fmt.Println("========== [ Hapus Mahasiswa ] ==========")

									fmt.Println("-[+] Error : id yang kamu cari tidak ditemukan!")

									fmt.Println("========== [ Hapus Mahasiswa ] ==========")

								}

							} else {

								fmt.Println("========== [ Hapus Mahasiswa ] ==========")

								fmt.Println("-[+] Error : Silakan masukan input yang tersedia!")

								fmt.Println("========== [ Hapus Mahasiswa ] ==========")

							}

						} else {

							fmt.Println("========== [ Hapus Mahasiswa ] ==========")

							fmt.Println("-[+] Warning : Id tidak ditemukan")

							fmt.Println("========== [ Hapus Mahasiswa ] ==========")

							sleep()

							continue

						}

					} else {

						fmt.Println("========== [ Hapus Mahasiswa ] ==========")

						fmt.Println("-[+] Warning : Belum ada data mahasiswa")

						fmt.Println("========== [ Hapus Mahasiswa ] ==========")

						break

					}

					continue

				case 2:

					clearScreen()

​

					for i := 0; i < len(mHs)-1; i++ {

						tmp = append(tmp, mHs[i])

					}

​

					mHs = tmp

​

					if createFiles(DATABASE, jsonEncode(mHs)) {

						fmt.Println("========== [ Hapus Mahasiswa ] ==========")

						fmt.Println("-[+] Success : hapus berhasil !")

						fmt.Println("========== [ Hapus Mahasiswa ] ==========")

					} else {

						fmt.Println("========== [ Hapus Mahasiswa ] ==========")

						fmt.Println("-[+] Error : hapus gagal!")

						fmt.Println("========== [ Hapus Mahasiswa ] ==========")

					}

					if tryAgain("Apakah anda ingin kembali kemenu utama? (Y/n) : ") {

						break

					} else {

						continue

					}

				case 3:

					break // untuk keluar dari switch case

				default:

					fmt.Println("========== [ Hapus Mahasiswa ] ==========")

					fmt.Println("-[+] Warning : Silakan pilih menu yang tersedia!")

					fmt.Println("========== [ Hapus Mahasiswa ] ==========")

					sleep()

					continue

				}

				break // untuk keluar dari function

			} else {

				fmt.Println("========== [ Hapus Mahasiswa ] ==========")

				fmt.Println("-[+] Warning : Silakan pilih menu yang tersedia!")

				fmt.Println("========== [ Hapus Mahasiswa ] ==========")

			}

			sleep()

		}

	} else {

		fmt.Println("========== [ Hapus Mahasiswa ] ==========")

		fmt.Println("-[+] Warning : Belum ada data mahasiswa")

		fmt.Println("========== [ Hapus Mahasiswa ] ==========")

	}

}

​

func viewMahasiswa() {

	var inputIndex int

	var err error

	for {

		clearScreen()

		if len(mHs) > 0 {

			fmt.Println("========== [ Lihat Mahasiswa ] ==========")

			fmt.Println("1. Lihat mahasiswa berdasarkan id\n2. Lihat Semua\n3. Back")

			fmt.Print("Masukan Pilihan: ")

			inputIndex, err = strconv.Atoi(inputData())

			if err == nil {

				switch inputIndex {

				case 1:

					var id int

					var dumpID string

					var err error

​

					for i := range mHs {

						dumpID += strconv.Itoa(i + 1)

					}

					fmt.Println("========== [ Lihat Mahasiswa By id ] ==========")

					fmt.Println("List ID : ", strings.Split(dumpID, ""))

					fmt.Print("Masukan id: ")

					id, err = strconv.Atoi(inputData())

					if err == nil {

						if id <= len(mHs) && id > 0 {

							fmt.Println("========== [ Lihat Mahasiswa By id ] ==========")

							fmt.Println("id: ", id)

							fmt.Println("Nama : ", mHs[id-1].Nama)

							fmt.Println("Umur : ", mHs[id-1].Umur)

							fmt.Println("Jurusan : ", mHs[id-1].Jurusan)

							if tryAgain("Apakah anda ingin kembali kemenu utama? (Y/n) : ") {

								break

							} else {

								continue

							}

						} else {

							fmt.Println("========== [ Lihat Mahasiswa ] ==========")

							fmt.Println("-[+] Warning : Id tidak ditemukan")

							fmt.Println("========== [ Lihat Mahasiswa ] ==========")

						}

					} else {

						fmt.Println("========== [ Lihat Mahasiswa ] ==========")

						fmt.Println("-[+] Error : input harus angka boss!")

						fmt.Println("========== [ Lihat Mahasiswa ] ==========")

					}

					sleep()

					continue

				case 2:

					fmt.Println("========== [ Lihat Semuah Mahasiswa ] ==========")

					for i := range mHs {

						fmt.Println("========== [ Lihat Semuah Mahasiswa ] ==========")

						fmt.Println("id: ", i+1)

						fmt.Println("Nama : ", mHs[i].Nama)

						fmt.Println("Umur : ", mHs[i].Umur)

						fmt.Println("Jurusan : ", mHs[i].Jurusan)

					}

					if tryAgain("Apakah anda ingin kembali kemenu utama? (Y/n) : ") {

						break

					} else {

						continue

					}

				case 3:

					break

				default:

					fmt.Println("========== [ Lihat Mahasiswa ] ==========")

					fmt.Println("-[+] Error : Silakan pilih menu yang tersedia!")
					fmt.Println("========== [ Lihat Mahasiswa ] ==========")
					sleep()
					continue
				}
				break
			} else {
				fmt.Println("========== [ Lihat Mahasiswa ] ==========")
				fmt.Println("-[+] Warning : Input harus berupa angka! ")
				fmt.Println("========== [ Lihat Mahasiswa ] ==========")
			}
		} else {
			fmt.Println("========== [ Lihat Mahasiswa ] ==========")
			fmt.Println("-[+] Warning : belum ada data Mahasiswa")
			fmt.Println("========== [ Lihat Mahasiswa ] ==========")
			break
		}
		sleep()
	}
}

​

func createFiles(path string, str string) bool {
	os.Remove(path)
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err == nil {
		defer file.Close()
		_, err = file.WriteString(str)
		if err == nil {
			err = file.Sync()
			if err == nil {
				return true
			}
			return false
		}
	}
	return false
}

func readFile(path string) string {
	var text string
	var file, _ = os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	getText := bufio.NewScanner(file)
	for getText.Scan() {
		text += getText.Text()
	}
	return text
}

func jsonEncode(data []mahasiswa) string {
	var jsonData, err = json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	var jsonString = string(jsonData)
	return jsonString
}

func jsonDecode(data string) []mahasiswa {

	var user []mahasiswa
	var jsonData = []byte(data)
	var err = json.Unmarshal(jsonData, &user)
	if err != nil {

		fmt.Println(err.Error())
	}

	// fmt.Println(user)

	// for i := range user {

	// 	user = append(user, mahasiswa{Id: user[i].Id, Nama: user[i].Nama, Umur: user[i].Umur, Jurusan: user[i].Jurusan})

	// }
	return user
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func inputData() string {

	var tmpInput string
	getInput := bufio.NewScanner(os.Stdin)
	getInput.Scan()
	tmpInput = getInput.Text()
	return tmpInput
}

func tryAgain(str string) bool {
	fmt.Print(str)
	sheAsk := strings.ToLower(inputData())
	if sheAsk == "y" {
		return true
	}
	return false
}

func typeOfVar(str interface{}) string {

	_, err := strconv.Atoi(str.(string))

	maybe := strings.ContainsAny(str.(string), "0123456789")

	if err == nil && maybe {

		return "int"
	} else if !maybe {

		return "string"
	}

	return "unknow"

}

func sleep() {

	time.Sleep(1 * time.Second)
}


func clearScreen() {

	c := exec.Command("clear")

	c.Stdout = os.Stdout

	c.Run()
}
