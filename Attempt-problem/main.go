package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scanln(&X, &Y) // Read X and Y

	unattempted := X - Y // Calculate the number of unattempted problems
	fmt.Println(unattempted)
}
