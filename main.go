package main

import "fmt"

func main() {

	var inputUser string

	fmt.Println("Masukkan apa saja: ")
	fmt.Scanln(&inputUser)

	hasil, status := search(inputUser)

	if status == true {
		fmt.Printf("Karakter %s tidak nyata mas", hasil)
	} else {
		fmt.Println("Karakter yang kau cari masih ada kemungkinan nyata")
	}

}

func search(chara string) (hasil string, status bool) {
	karakter := [...]string{"Emilia", "Saki"}

	for _, karak := range karakter {
		if chara == karak {
			hasil = chara
			status = true
			return
		}
	}
	hasil = chara
	status = false
	return
}
