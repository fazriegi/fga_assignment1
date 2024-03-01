package main

import (
	"fmt"
	"os"
)

type Student struct {
	Absen     string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	students := []Student{
		{"1", "Fazri", "Bogor", "Mahasiswa", "Menambah pengalaman dan relasi"},
		{"2", "Egi", "Jakarta", "Mahasiswa", "Karna maskotnya lucu"},
		{"3", "Ramadhan", "Bekasi", "Backend Developer", "Mengisi waktu luang"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <no-absen>")
		return
	}

	absen := os.Args[1]
	student := getStudentByAbsen(absen, students)

	if student.Absen == "" {
		fmt.Println("Data tidak ditemukan")
		return
	}

	fmt.Println("Data Mahasiswa:")
	fmt.Println("===============")
	fmt.Println("Nama:", student.Nama)
	fmt.Println("Alamat:", student.Alamat)
	fmt.Println("Pekerjaan:", student.Pekerjaan)
	fmt.Println("Alasan:", student.Alasan)
}

func getStudentByAbsen(absen string, students []Student) Student {
	for _, v := range students {
		if v.Absen == absen {
			return v
		}
	}
	return Student{}
}
