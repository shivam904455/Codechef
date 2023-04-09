package main

import "fmt"

func main() {
    // Create a slice to hold the input pairs
    var pairs [][]int

    // Add the input pairs to the slice
    pairs = append(pairs, []int{1, 4})
    pairs = append(pairs, []int{3, 4})
    pairs = append(pairs, []int{4, 2})
    pairs = append(pairs, []int{2, 6})

    // Check whether the second number in each pair is divisible by the first 
    for _, pair := range pairs {
        if pair[1]%pair[0] == 0 {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
