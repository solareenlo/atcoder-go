package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n int
	a = [200005]int{}
	b = [200005]int{}
)

func P(x int) int {
	for i := 1; i < n*2; i++ {
		if a[i] >= x {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
	b[0] = b[1]
	for i := n; i <= n+n; i++ {
		if b[i] == b[i+1] {
			return b[i]
		}
		if b[n+n-i] == b[n+n-i-1] {
			return b[n+n-i]
		}
	}
	return 0
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i < n*2; i++ {
		fmt.Fscan(in, &a[i])
	}

	l := 0
	r := 1 << 60
	ans := 0
	for l <= r {
		mid := (l + r) / 2
		if P(mid) != 0 {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}
	}
	fmt.Println(ans)
}
