// Dup1 will show the content of each line that appears more than once
//when you type input at the terminal and shows a counter.
// run : go run dup1.go  + ENTER + write your text and when you are done
// press CTRL + Z (tried on Windows with git bash terminal)

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// A map holds key/value pair - keys and values can be of any type.
	// Here keys are of type String and values of type int.
	//The make() method, creates a new empty map.
	counter := make(map[string]int)
	// the program read the argument(s) at command line,
	// the line is the key into a map and the corresponding
	// value is incremented
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counter[input.Text()]++
	}
	// handeling eventual errors from input.Err()
	input.Err()
	// contrary to other languages, we don't put
	// the arguments of the for loop in parentheses but braces are required for the body.
	for line, n := range counter {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
