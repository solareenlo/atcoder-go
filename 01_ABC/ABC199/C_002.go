package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, q int
	var s []byte
	fmt.Scan(&n, &s, &q)
	flip := false
	for i := 0; i < q; i++ {
		var t, a, b int
		fmt.Fscan(in, &t, &a, &b)
		a--
		b--
		if t == 1 {
			if flip {
				a, b = (a+n)%(n*2), (b+n)%(n*2)
			}
			s[a], s[b] = s[b], s[a]
		} else {
			flip = !flip
		}
	}
	if flip {
		fmt.Println(string(s[n:]) + string(s[:n]))
	} else {
		fmt.Println(string(s))
	}
}
