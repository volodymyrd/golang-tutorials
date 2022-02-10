// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
//
// go run ch1/dup/dup1/dup1.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		if s == "-1" {
			break
		}
		counts[s]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
