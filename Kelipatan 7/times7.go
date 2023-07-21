package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	if angka%7 == 0 {
		fmt.Printf("%d adalah kelipatan dari 7.\n", angka)
	} else {
		fmt.Printf("%d bukan kelipatan dari 7.\n", angka)
	}
}
