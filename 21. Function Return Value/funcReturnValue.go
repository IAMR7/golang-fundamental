package main

import "fmt"

func main() {

	//pada function main kita hanya perlu memanggil function getHello beserta parameternya
	result := getHello("Reza Febriansyah")
	fmt.Println(result)
}

//function memiliki parameter yag mengharuskan memasukan string serta mengembalikan data string
func getHello(name string) string {

	//string dikembalikan dengan return
	return "Hello, " + name
}
