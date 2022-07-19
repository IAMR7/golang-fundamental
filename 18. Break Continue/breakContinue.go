package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		//misal cek apakah angka habis dibagi 2, jika iya continue/skip aja
		if i%2 == 0 {
			continue
		}
		//makanya ter print lah ini, yang tampil cuman angka perulangan yang tidak habis dibagi 2
		fmt.Println("Perulangan ke ", i)
	}

}
