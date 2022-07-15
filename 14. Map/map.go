package main

import "fmt"

func main() {

	person := map[string]string{

		"nama":   "Reza",
		"alamat": "Jakarta",
	}

	//misal tambah data baru di map
	person["pekerjaan"] = "UI UX Engineer"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person["pekerjaan"])

	//cari tau jumlah datanya di map
	fmt.Println(len(person))

	//misal bikin map baru dari 0 banget
	buku := make(map[string]string)
	buku["judul"] = "Buku GOLANG"
	buku["penulis"] = "Reza"
	buku["salah"] = "Ups"

	delete(buku, "salah")
	fmt.Println(buku)
}
