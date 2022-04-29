package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	type pair struct {
		x string
		y int
	}
	S := make([]pair, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		m := len(s)
		for j := 0; j < m+1; j++ {
			if j == m || s[j] != '0' {
				S[i] = pair{s[j:], j}
				break
			}
		}
	}

	sort.Slice(S, func(i, j int) bool {
		if len(S[i].x) != len(S[j].x) {
			return len(S[i].x) < len(S[j].x)
		}
		if S[i].x != S[j].x {
			return S[i].x < S[j].x
		}
		return S[i].y > S[j].y
	})

	for i := 0; i < n; i++ {
		fmt.Fprint(out, strings.Repeat("0", S[i].y))
		fmt.Fprintln(out, S[i].x)
	}
}
