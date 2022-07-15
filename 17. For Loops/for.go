package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		fmt.Println("Ini adalah perulangan ke", i)

	}

	slice := []string{"Reza", "Akbar", "Putri", "Imam"}

	for i := 0; i < len(slice); i++ {
		fmt.Println("Temanku", slice[i])
	}

	//For Range
	for _, value := range slice {
		// fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}

}
