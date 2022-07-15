package main

import "fmt"

func main() {

	nama := "Reza"

	switch nama {
	case "Reza":
		fmt.Println("Halo Reza")
	case "Febri":
		fmt.Println("Halo Febri")
	case "Rian":
		fmt.Println("Halo Rian")
	default:
		fmt.Println("Halo terserah siapa aja")
	}

	//short statement dari Switch Expression
	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama cukup")
	}

	//Switch tanpa kondisi
	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("Nama kepanjangan")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
