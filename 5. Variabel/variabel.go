package main

import "fmt"

func main() {

	//declare variabel 1 line
	var saya string = "Reza Febriansyah"
	fmt.Println(saya)

	//tanpa menggunakan var dengan (:=)
	umur := 22
	fmt.Println(umur)

	//multiple variabel
	var (
		namadepan    = "Reza"
		namabelakang = "Febriansyah"
	)

	fmt.Print(namadepan)
	fmt.Print(" ")
	fmt.Print(namabelakang)
}
