package main

import (
	"fmt"
)

func main() {
	// Slice
	numbers := []int{1, 2, 3, 4}

	// Maps
	person := map[string]string{
		"first_name":   "Poorva",
		"last_name":    "Mahajan",
		"company_name": "Oogway",
	}

	fmt.Println("First two elements", numbers[:2])
	fmt.Println("First Name:", person["first_name"])
}
