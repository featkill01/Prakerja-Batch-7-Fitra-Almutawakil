package main

import "fmt"

func hitungLuasTrapesium(alas, atas, tinggi float64) float64 {
	return 0.5 * tinggi * (alas + atas)
}

func main() {
	var alas, atas, tinggi float64
	fmt.Print("Masukkan panjang alas: ")
	fmt.Scan(&alas)

	fmt.Print("Masukkan panjang atas: ")
	fmt.Scan(&atas)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scan(&tinggi)

	luas := hitungLuasTrapesium(alas, atas, tinggi)
	fmt.Println("Luas trapesium:", luas)
}
