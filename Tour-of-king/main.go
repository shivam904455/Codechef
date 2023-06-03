package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var N, M int
		fmt.Scan(&N, &M)

		totalPeople := N*5 + M*7
		fmt.Println(totalPeople)
	}
}

