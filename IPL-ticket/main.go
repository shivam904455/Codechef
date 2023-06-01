package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)
  
    for i := 0; i < T; i++ {
        var N, M int
        fmt.Scan(&N, &M)
      
        if N <= M {
            fmt.Println(0)
        } else {
            fmt.Println(N - M)
        }
    }
}
