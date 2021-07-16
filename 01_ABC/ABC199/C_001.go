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
	flip := 0
	for i := 0; i < q; i++ {
		var t, a, b int
		fmt.Fscan(in, &t, &a, &b)
		a--
		b--
		if t == 1 {
			s[(a+n*flip)%(n*2)], s[(b+n*flip)%(n*2)] = s[(b+n*flip)%(n*2)], s[(a+n*flip)%(n*2)]
		} else {
			if flip == 0 {
				flip = 1
			} else {
				flip = 0
			}
		}
	}
	if flip == 1 {
		fmt.Println(string(s[n:]) + string(s[:n]))
	} else {
		fmt.Println(string(s))
	}
}
