package main

import "fmt"

func main() {
	var T int
	fmt.Scanln(&T) // Read the number of test cases

	for i := 0; i < T; i++ {
		var X, Y int
		fmt.Scanln(&X, &Y) // Read X and Y for each test case

		total := X * Y // Calculate the total amount of money Chef will have to pay
		fmt.Println(total) // Print the result
	}
}
