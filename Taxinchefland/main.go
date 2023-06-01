package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var X int
		fmt.Scan(&X)

		if X > 100 {
			X -= 10
		}

		fmt.Println(X)
	}
}
