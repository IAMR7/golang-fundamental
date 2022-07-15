package main

import "fmt"

func main() {

	var nilaiakhir = 90
	var absensi = 80

	var lulusnilaiakhir = nilaiakhir > 85
	var lulusabsensi = absensi > 80

	var lulus = lulusnilaiakhir && lulusabsensi

	fmt.Println(lulus)
}
