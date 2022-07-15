package main

import "fmt"

func main() {
	const nama1 = "Reza"
	const nama2 = "Akbar"
	const hasil = nama1 == nama2
	fmt.Println(hasil)

	const angka1 = 100
	const angka2 = 150
	const hasilangka bool = angka1 < angka2
	fmt.Println(hasilangka)
}
