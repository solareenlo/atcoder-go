package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)

	mo := make(map[string]bool)
	for q > 0 {
		q--
		var op int
		var a, b string
		fmt.Fscan(in, &op, &a, &b)
		switch op {
		case 1:
			mo[a+" "+b] = true
		case 2:
			mo[a+" "+b] = false
		case 3:
			if mo[a+" "+b] && mo[b+" "+a] {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
