package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s, t string
	fmt.Fscan(in, &s, &t)

	sa := make([]int, 233333)
	for i := 0; i < len(s); i++ {
		sa[i+1] = (sa[i] + int(s[i])%4) % 3
	}
	sb := make([]int, 233333)
	for i := 0; i < len(t); i++ {
		sb[i+1] = (sb[i] + int(t[i])%4) % 3
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		if (sa[b]-sa[a-1]-sb[d]+sb[c-1])%3 != 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
