package main

import "fmt"

func main() {

	const a = 5
	const b = 8
	var hasil = a + b
	fmt.Println(hasil)

	var i = 10
	//mempersingkat sintaks, sama saja dengan i = i + 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

}
