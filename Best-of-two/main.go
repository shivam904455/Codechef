package main
import (
  "fmt"
  "math"
)

func main(){
	var cases int
	var x, y float64
	fmt.Scanln(&cases)
	
	for i:=0; i < cases; i++ {
	  fmt.Scanln(&x, &y)
    fmt.Println(math.Max(x, y))	  
	}
}