package main

import "fmt"

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)

	rating := N - A
	ratingGreaterThan1000 := rating - B

	fmt.Println(rating, ratingGreaterThan1000)
}
