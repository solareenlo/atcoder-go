package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	r, b, a, x := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for i := 1; i < len(s); i++ {
			if s[i-1] == 'A' && s[i] == 'B' {
				r++
			}
		}
		if s[0] == 'B' {
			b++
		}
		if s[len(s)-1] == 'A' {
			a++
		}
		if s[0] == 'B' && s[len(s)-1] == 'A' {
			x++
		}
	}

	if a != 0 && a == x && b == x {
		r--
	}
	fmt.Println(r + min(a, b))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
