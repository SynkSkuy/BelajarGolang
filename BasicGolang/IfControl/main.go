package main

import "fmt"

func main() {
	score := 75
	// Your code here
	var grade string
	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Println("Score:", score, "Grade:", grade)
}
