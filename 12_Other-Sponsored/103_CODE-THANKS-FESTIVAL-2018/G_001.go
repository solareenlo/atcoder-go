package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	var a, b, ra [5005]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		ra[a[i]] = i
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
		b[i]--
	}
	ans := n * (n + 1) / 2
	for c := 0; c < n-k; c++ {
		var used [5005]bool
		dl, dr := -1, -1
		for i := 0; i < n; i++ {
			if used[i] || ra[i] == -1 {
				continue
			}
			x, r := i, i
			for !used[x] && ra[x] != -1 {
				used[x] = true
				x = b[ra[x]]
				r = max(r, x)
			}
			if r-i > dr-dl {
				dr = r
				dl = i
			}
		}
		if dr == -1 {
			continue
		}
		x := dl
		last := -1
		for x != dr {
			j := ra[x]
			ra[x] = last
			a[j], b[j] = b[j], a[j]
			x = a[j]
			last = j
		}
		ans += dr - dl
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
