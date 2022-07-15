package main

import "fmt"

func main() {
	var nilai32 int32 = 1000000000
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	//konversi dari byte ke string "huruf"
	var nama string = "Reza Febriansyah"
	var e = nama[7]
	var estring string = string(e)

	fmt.Println(nama)
	fmt.Println(estring)
}
