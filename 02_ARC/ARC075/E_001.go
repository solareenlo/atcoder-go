package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	a   = [200002]int{}
	b   = [200002]int{}
	ans int
)

func sort(l, r int) {
	if l == r {
		if a[l] >= 0 {
			ans++
		}
		return
	}
	m := (l + r) >> 1
	sort(l, m)
	sort(m+1, r)
	for i, p, j := l, l, m+1; p <= r; p++ {
		if (i == m+1 || a[i] > a[j]) && j != r+1 {
			b[p] = a[j]
			j++
		} else {
			ans += r - j + 1
			b[p] = a[i]
			i++
		}
	}
	for i := l; i <= r; i++ {
		a[i] = b[i]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[i] = a[i-1] + x - k
	}
	sort(1, n)
	fmt.Println(ans)
}
