package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)
  
    for i := 0; i < T; i++ {
        var X, Y int
        fmt.Scan(&X, &Y)
      
        hours := (Y - X + 12) % 12
      
        fmt.Println(hours)
    }
}
