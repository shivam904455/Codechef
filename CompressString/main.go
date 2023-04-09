package main

import "fmt"

func main() {
	// input:=aaabbcaa
	// output:=3a2b1c2a
	s := "aaabbcaa"
	run := []rune(s)
	fresult := ""
	count := 1
	for i := 0; i < len(run); i++ {
		// fmt.Println("rune = ", string(run[1]), "count =",count , "i =",i,"len(run)=",len(run))
		result := ""
		if i+1 >= len(run) {
			result = result + fmt.Sprintf("%d", count) + string(run[1])
			fresult = fresult + result
			// fmt.Println(result)
			break
		}
		if run[i] == run[i+1] {
			count++
		} else {
			// fmt.Println("rune---=",string(run[i]))
			result = fmt.Sprintf("%d", count) + string(run[i])
			count = 1
			fresult = fresult + result
			// fmt.Println(reslut)
		}
	}
	fmt.Println(fresult)

	

}
