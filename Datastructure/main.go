package main

import "fmt"

type Student struct {
	Name       string
	Class      int
	Rollnumber int
}

func main() {

	 var shivam Student
		shivam.Name =" Shivam singh"
shivam.Class =12
shivam.Rollnumber=33
	 
	Aashu := Student{
		Name:       "Aashu kumer",
		Class:      12,
		Rollnumber: 23,
	}

	fmt.Println("this is a vishal sturct", Aashu)
	fmt.Println("this is a vishal sturct", shivam)
}