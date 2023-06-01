package main

import "fmt"

func main() {
	var T int
	fmt.Scanln(&T) // Read the number of test cases

	for i := 0; i < T; i++ {
		var A, B int
		fmt.Scanln(&A, &B) // Read A and B for each test case

		maxBurgers := min(A, B) // Find the minimum of A and B, as that will limit the number of burgers Chef can make
		fmt.Println(maxBurgers)
	}
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
