package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var testCases int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fscan(reader, &testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var a int
		fmt.Fscan(reader, &a)

		fmt.Fprintln(writer, a*15)
	}
}
