package main

import "fmt"

func main() {
	length := 5

	for i := 1; i <= length; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("a") // Print "a" for each iteration
		}
		fmt.Println() // Move to the next line after each inner loop iteration
	}
}
