package main

import "fmt"

// Mendefinisikan konstanta untuk angka Romawi
const (
	I   = 1
	X   = 10
	VII = 7
	L   = 50
	C   = 100
)

func main() {

	// Menggunakan for loop untuk menampilkan angka dari I hingga VII
	for i := I; i <= VII; i++ {
		fmt.Println("Saya sedang belajar bahasa pemrograman Go!", i)
	}

	fmt.Println("=====================================")

	// Menggunakan while loop dengan for
	i := I
	for i <= VII {
		fmt.Println("Saya sedang belajar bahasa pemrograman Go!", i)
		i++

	}

	fmt.Println("=====================================")

	title := "Belajar Bahasa Pemrograman Go"

	// Menggunakan fmt.Printf untuk menampilkan karakter dalam bentuk string
	for index, char := range title {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	fmt.Println("=====================================")
	//Menggunakan fmt.Printf untuk menampilkan karakter dalam bentuk integer (ASCII/Unicode)
	for index, char := range title {
		fmt.Printf("Index: %d, Character: %d\n", index, char)
	}
}
