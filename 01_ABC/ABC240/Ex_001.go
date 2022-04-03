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
	s = " " + s

	type pair struct{ x, y int }
	V := make([][]pair, 25010)
	f := make([]string, 0)
	for i := 1; i <= n; i++ {
		var t string
		for j, k := i, 0; j <= min(n, i+250); j++ {
			t += string(s[j])
			for k < len(f) && f[k] < t {
				k++
			}
			V[j] = append(V[j], pair{k, i})
			if k == len(f) {
				break
			}
		}
		for _, p := range V[i] {
			t = s[p.y : i+1]
			if p.x == len(f) {
				f = append(f, t)
			} else {
				f[p.x] = minStr(f[p.x], t)
			}
		}
	}

	fmt.Println(len(f))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minStr(a, b string) string {
	if a < b {
		return a
	}
	return b
}
