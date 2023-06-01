package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)
  
    for i := 0; i < T; i++ {
        var X, N int
        fmt.Scan(&X, &N)
      
        pointsPerTest := X / 10
        score := pointsPerTest * N
      
        fmt.Println(score)
    }
}
