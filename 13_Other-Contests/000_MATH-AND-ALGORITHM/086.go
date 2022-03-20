package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	ans := "Yes"
	a := 0
	for _, c := range s {
		if c == '(' {
			a++
		}
		if c == ')' {
			a--
		}
		if a < 0 {
			ans = "No"
		}
	}
	fmt.Println(ans)
}
