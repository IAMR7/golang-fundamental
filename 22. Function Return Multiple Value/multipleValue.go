package main

import "fmt"

func main() {

	//cara normal

	// firstName, lastName := getFullName()
	// fmt.Println("Hallo,", firstName, lastName)

	//cara meng-ignore return value

	firstName, _ := getFullName()
	fmt.Println("Hallo,", firstName)
}

func getFullName() (string, string) {
	return "Reza", "Febriansyah"
}
