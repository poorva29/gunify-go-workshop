package main

import (
	"fmt"
)

// Add two float64

// Concatinate two strings
func concatinate(str1, str2 string) string {
	return str1 + " " + str2
}

// Add two float32
func add32(a, b float32) (res float32) {
	res = a + b
	return
}

func main() {
	a, b := 4.5, 5.5
	fmt.Printf("\nAddition of add for %f and %f is %f\n", a, b, add(a, b))

	first_name, last_name := "Poorva", "Mahajan"
	fmt.Printf("\nConcatination of %s and %s is %s\n", first_name, last_name, concatinate(first_name, last_name))

	fmt.Printf("\nResult of add32 for %f and %f is %f\n", a, b, add32(a, b))
}
