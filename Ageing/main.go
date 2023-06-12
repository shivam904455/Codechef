package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var X int
		fmt.Scan(&X)

		chefAge := 20
		chefinaAge := 10 + (X - chefAge)

		fmt.Println(chefinaAge)
	}
}
 