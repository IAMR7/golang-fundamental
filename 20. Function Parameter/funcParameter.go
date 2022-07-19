package main

import "fmt"

func main() {
	sayHelloTo("Reza", "Febriansyah")
}

//urutan dari pendefinisian varibel parameter harus sesuai urutannya
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello,", firstName, lastName)
}
