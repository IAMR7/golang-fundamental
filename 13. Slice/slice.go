package main

import "fmt"

func main() {
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice = bulan[4:7]

	fmt.Println(slice)
	fmt.Println(len(slice)) //panjang slice
	fmt.Println(cap(slice)) //capacity slice

	slice[0] = "Diubah"
	fmt.Println(bulan) //Jika data slice ada yang diubah maka yang di array juga berubah

	var slice1 = bulan[11:]
	fmt.Println(slice1)

	var slice2 = append(slice1, "Reza")
	fmt.Println(slice2)

	slice2[1] = "Bukan Desember"
	fmt.Println(slice2)

	fmt.Println(bulan)

	//biki slice baru dengan array yang tersembunyi
	newSLice := make([]string, 2, 5)

	newSLice[0] = "Reza"
	newSLice[1] = "Febriansyah"

	fmt.Println(newSLice)
	fmt.Println(len(newSLice))
	fmt.Println(cap(newSLice))

	//menyalin/copy slice
	copySlice := make([]string, len(newSLice), cap(newSLice))
	copy(copySlice, newSLice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
