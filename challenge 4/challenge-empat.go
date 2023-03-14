package main

import (
	"fmt"
	"os"
	"strconv"
)

// Deklarasi struct Teman dengan beberapa atribut seperti Nama, Alamat, Pekerjaan, dan Alasan.
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Inisialisasi slice temanList dengan beberapa data teman.
var temanList = []Teman{
	{"Agung", "Jl. Moh. Kahfi 1 No.19", "Developer", "Ingin belajar bahasa pemrograman baru"},
	{"Chintia", "Jl. Moh. Kahfi 1 No.11", "Cloud Architect", "Ingin mengembangkan skill programming"},
	{"Jeavira", "Jl. Moh. Kahfi 1 No.99", "IT Support", "Ingin belajar bahasa pemrograman yang lebih cepat dan efisien"},
}

func main() {
	args := os.Args // Mengambil argumen yang diinputkan pada saat menjalankan program.
	// Jika jumlah argumen yang diinputkan bukan 2, maka akan menampilkan pesan error dan program akan berhenti.
	if len(args) != 2 {
		fmt.Println("Usage: go run challenge-empat.go <nomor absen>")
		return
	}

	// Konversi argumen ke integer dan cek apakah argumen berupa angka atau tidak.
	index, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	// Cek apakah nomor absen ada di antara 1 dan panjang slice temanList.
	if index < 1 || index > len(temanList) {
		fmt.Println("Tidak ada teman dengan nomor absen tersebut")
		return
	}

	// Menampilkan data teman dengan nomor absen yang diinputkan.
	teman := temanList[index-1]
	fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas Golang: %s\n",
		teman.Nama, teman.Alamat, teman.Pekerjaan, teman.Alasan)
}
