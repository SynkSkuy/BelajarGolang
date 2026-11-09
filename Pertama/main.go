package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Hello, What'up Bro")

	result := calculation.Add(100, 300)

	fmt.Println("Hasil dari penjumlahan adalah", result)

	result2 := calculation.Multiply(10, 20)
	fmt.Println("Hasil dari perkalian adalah", result2)
}
