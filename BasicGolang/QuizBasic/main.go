package main

import "fmt"

func main() {

	title := "Golang the best programming language"

	// disini kita akan menampilkan huruf yang berada di index genap
	for index, char := range title {

		if index%2 == 0 {
			fmt.Println("index:", index, "char:", string(char))
		}

	}

	fmt.Println("=====================================")

	// disini kita akan menampilkan huruf vokal yang ada di dalam string title
	for index, char := range title {

		letterString := string(char)

		switch letterString {
		case "a", "e", "i", "o", "u":
			fmt.Println("index:", index, "char:", letterString, "adalah huruf vokal")
		}
	}

}
