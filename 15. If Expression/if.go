package main

import "fmt"

func main() {

	nama := "Reza"

	//cek kondisi nya, true atau false
	if nama == "Reza" {
		fmt.Println("Beneeer")
	} else if nama == "Eja" {
		fmt.Println("yaudah boleh aja si")
	} else {
		fmt.Println("salah cok")
	}

	//short statement
	if panjang := len(nama); panjang > 5 {
		fmt.Println("Kepanjangan namanya")
	} else {
		fmt.Println("nama udah cukup")
	}

}
