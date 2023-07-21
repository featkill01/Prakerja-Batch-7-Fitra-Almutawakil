package main

import "fmt"

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println(n, "adalah bilangan prima.")
	} else {
		fmt.Println(n, "bukan bilangan prima.")
	}
}
