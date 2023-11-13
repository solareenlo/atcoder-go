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

	var x string
	fmt.Fscan(in, &x)

	var p [26]int
	for j := 0; j < 5; j++ {
		p[x[j]-'A']++
	}

	r, u := -1, -1
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)

		var q [26]int
		for j := 0; j < 4; j++ {
			q[s[j]-'A']++
		}

		t := 0
		for c := 0; c < 26; c++ {
			t = max(t, (min(q[c], 2)+min(p[c], 3)+1)*26-c)
		}
		if t > u {
			u = t
			r = i
		}
	}
	fmt.Println(r + 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
