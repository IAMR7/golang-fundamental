package main

import "fmt"

func main() {

	//bikin alias dari tipedata yang sudah ada
	type Noktp string
	type Nikah bool

	var ktpku Noktp = "121348489101121983"
	var akunikah Nikah = false

	fmt.Println(ktpku)
	fmt.Println(akunikah)
}
