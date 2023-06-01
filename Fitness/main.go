package main

import "fmt"

func main() {
	var T int
	fmt.Scanln(&T) // Read the number of test cases

	for i := 0; i < T; i++ {
		var X int
		fmt.Scanln(&X) // Read X for each test case

		total := X * 2 * 5 // Calculate the total distance Chef travels in a week (2 trips per day, 5 working days)
		fmt.Println(total) // Print the result
	}
}
