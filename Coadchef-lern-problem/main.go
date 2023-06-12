package main

import "fmt"

func totalCoursesInSection(N int) int {
	return 2 * N
}

func main() {
	var N int
	fmt.Scanln(&N)

	totalCourses := totalCoursesInSection(N)
	fmt.Println(totalCourses)
}
