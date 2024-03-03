package main

import (
	"fmt"
	"os"
)

// Teman struct untuk menyimpan data teman
type BiodataTeman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (biodata BiodataTeman) LihatAlasan() string {
	return " Memilih kelas Golang karena " + biodata.Alasan
}

var Biodata = map[int]BiodataTeman{
	1: {"Khairil", "Jl. NNN", "Mahasiswa", "Ingin menambah portofolio pengalaman"},
	2: {"Kevin", "Jl. NNN", "Mahasiswa", "Menambah skill Golang"},
	3: {"Arifin", "Jl. NNN", "Mahasiswa", "Sebagai Backend Engineer"},
	4: {"Anugrah", "Jl. NNN", "Mahasiswa", "Sebagai pemula Backend Golang Engineer"},
	5: {"Kurniawan", "Jl. NNN", "Mahasiswa", "Ingin menambah relasi dengan mentor"},
}

func ShowBiodata(absen int) {
	fmt.Printf("Data Teman dengan Absen %d:\n", absen)
	fmt.Printf("Nama: %s\n", Biodata[absen].Nama)
	fmt.Printf("Alamat: %s\n", Biodata[absen].Alamat)
	fmt.Printf("Pekerjaan: %s\n", Biodata[absen].Pekerjaan)
	fmt.Printf("Alasan: %s\n", Biodata[absen].LihatAlasan())
}

func main() {
	var absenInt int

	if len(os.Args) != 2 {
		fmt.Println("Gunakan command: go run biodata.go <nomor_absen>")
		os.Exit(1)
	}

	if _, err := fmt.Sscanf(os.Args[1], "%d", &absenInt); err != nil {
		fmt.Println("Tidak ")
		os.Exit(1)
	}

	ShowBiodata(absenInt)
}
