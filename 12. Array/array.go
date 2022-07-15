package main

import "fmt"

func main() {

	//berdasarkan index
	var nama [3]string
	nama[0] = "Reza"
	nama[1] = "Akbar"
	nama[2] = "Putri"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var value = [3]int{
		50,
		20,
		35,
	}
	fmt.Println(value)
}
