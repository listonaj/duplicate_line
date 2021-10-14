// Dup1 will show the content of each line that appears more than once
//when you type input at the terminal and shows a counter.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counter := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counter[input.Text()]++
	}
	// handeling eventual errors from input.Err()
input.Err()
	for line, n := range counter {
		if n>1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}