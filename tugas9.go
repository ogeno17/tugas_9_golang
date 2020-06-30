package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer pulihkan()

	var input string
	fmt.Print("Masukan angka: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err != nil {
		fmt.Println("Ini bukan angka")
		panic(err.Error())
	} else {
		fmt.Println("ini adalah angka ", number)
	}
}

func pulihkan() {
	if r := recover(); r != nil {
		fmt.Println("Pulihkan!")
	} else {
		fmt.Println("Program Sukses!")
	}
}

/*
Analisis

baris 17 : ketika di convert, fungsi Atoi() mengembalikan nilai hasil dan error
baris 19 : jika terdapat error -> jalankan panic dan stop program
baris 9  : defer adalah sintaks yang dijalankan terkahir sebelum fungsi selesai/berhenti
baris 28 : seleksi kondisi jika recover terdapat error -> print "Pulihkan!"
baris 30 : seleksi kondisi jika recover tidak terdapat error -> print "Program Sukses!"

*/
