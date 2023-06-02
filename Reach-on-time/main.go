package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var X int
		fmt.Scan(&X)

		if X >= 30 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
