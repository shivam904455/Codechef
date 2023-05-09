package main

import "fmt"

func main() {

	var pairs [][]int

	pairs = append(pairs, []int{1, 4})
	pairs = append(pairs, []int{3, 4})
	pairs = append(pairs, []int{4, 2})
	pairs = append(pairs, []int{2, 6})

	for _, pair := range pairs {
		if pair[1]%pair[0] == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
