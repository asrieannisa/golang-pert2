package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	ID        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Tolong masukan nama atau nomor absen")
		fmt.Println("Contoh: 'go run main.go Fitri' atau 'go run main.go 2'")
		return
	}

	args := os.Args[1]

	if num, err := strconv.Atoi(args); err == nil {
		sliceToStructInt(num)
	} else {
		sliceToStructString(args)
	}

}

func sliceToStructInt(args int) {
	var people = []Biodata{
		{ID: 0, nama: "Asrie", alamat: "Bintaro", pekerjaan: "Software Maintenance", alasan: "Upgrade diri"},
		{ID: 1, nama: "Annisa", alamat: "Cikampek", pekerjaan: "Impelementor", alasan: "Tambah ilmu"},
		{ID: 2, nama: "Pratiwi", alamat: "Sukabumi", pekerjaan: "Developer", alasan: "Daripada gabut"},
		{ID: 3, nama: "Laura", alamat: "Jakarta", pekerjaan: "QA", alasan: "Tambah Networking"},
	}

	found := false

	for _, value := range people {
		if value.ID == args {
			fmt.Println("ID: ", value.ID)
			fmt.Println("Nama: ", value.nama)
			fmt.Println("Alamat: ", value.alamat)
			fmt.Println("Pekerjaan: ", value.pekerjaan)
			fmt.Println("Alasan: ", value.alasan)
			found = true
			break // Hentikan iterasi karena data sudah ditemukan
		}
	}

	if !found {
		fmt.Println("Tidak ada absen dengan nomor yang Anda input, silahkan input kembali")
	}
}

func sliceToStructString(args string) {
	var people = []Biodata{
		{ID: 0, nama: "Asrie", alamat: "Bintaro", pekerjaan: "Software Maintenance", alasan: "Upgrade diri"},
		{ID: 1, nama: "Annisa", alamat: "Cikampek", pekerjaan: "Impelementor", alasan: "Tambah ilmu"},
		{ID: 2, nama: "Pratiwi", alamat: "Sukabumi", pekerjaan: "Developer", alasan: "Daripada gabut"},
		{ID: 3, nama: "Laura", alamat: "Jakarta", pekerjaan: "QA", alasan: "Tambah Networking"},
	}

	found := false

	for _, value := range people {
		if value.nama == args {
			fmt.Println("ID: ", value.ID)
			fmt.Println("Nama: ", value.nama)
			fmt.Println("Alamat: ", value.alamat)
			fmt.Println("Pekerjaan: ", value.pekerjaan)
			fmt.Println("Alasan: ", value.alasan)
			found = true
			break // Hentikan iterasi karena data sudah ditemukan
		}
	}

	if !found {
		fmt.Println("Tidak ada absen dengan nama yang Anda input, silahkan input kembali")
	}
}
