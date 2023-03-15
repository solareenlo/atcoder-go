package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var k int
	fmt.Fscan(in, &k)

	var t string
	fmt.Fscan(in, &t)

	n := len(t)
	suf := make([]int, n+1)
	nxt := make([]int, n+1)
	for i := range nxt {
		nxt[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1]
		if t[i] == 's' {
			suf[i]++
		}
		if t[i] != 's' {
			nxt[i] = i
		} else {
			nxt[i] = nxt[i+1]
		}
	}

	var s string
	for i := 0; i < n; i++ {
		if t[i] != 's' || k == 0 {
			s += string(t[i])
		} else if suf[i] == k {
			k--
		} else if nxt[i] == n || t[nxt[i]] < 's' {
			k--
		} else {
			s += string(t[i])
		}
	}

	fmt.Fprintln(out, s)
}
