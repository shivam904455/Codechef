package main

import "fmt"

func isOnefulPair(a, b int) bool {
	sum := a + b + (a * b)
	return sum == 111
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	if isOnefulPair(a, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
