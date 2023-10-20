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
	s := make(map[int]bool)
	s[1] = true
	s[2] = true
	var r [100001]int
	for i := 1; i <= n; i++ {
		r[i] = i
	}
	c := 1
	for q > 0 {
		q--
		var a, b int
		fmt.Fscan(in, &a, &b)
		if a == c {
			c = b
		} else if b == c {
			c = a
		}
		r[a], r[b] = r[b], r[a]
		if c > 1 {
			s[r[c-1]] = true
		}
		if c < n {
			s[r[c+1]] = true
		}
	}
	fmt.Println(len(s))
}
