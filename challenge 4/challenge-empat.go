package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var temanList = []Teman{
	{"Agung", "Jl. Moh. Kahfi 1 No.19", "Developer", "Ingin belajar bahasa pemrograman baru"},
	{"Chintia", "Jl. Moh. Kahfi 1 No.11", "Cloud Architect", "Ingin mengembangkan skill programming"},
	{"Jeavira", "Jl. Moh. Kahfi 1 No.99", "IT Support", "Ingin belajar bahasa pemrograman yang lebih cepat dan efisien"},
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run challenge-empat.go <nomor absen>")
		return
	}

	index, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	if index < 1 || index > len(temanList) {
		fmt.Println("Tidak ada teman dengan nomor absen tersebut")
		return
	}

	teman := temanList[index-1]
	fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas Golang: %s\n",
		teman.Nama, teman.Alamat, teman.Pekerjaan, teman.Alasan)
}
