package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, s int
	fmt.Fscan(in, &n, &s)

	l := 1
	r := n
	ct := n
	x := make([]int, n+1)
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &p[i])
	}

	tp := make([]int, 100005)
	for l < r {
		if s < x[l] {
			p[l] = 1 << 60
		}
		if s > x[r] {
			p[r] = 1 << 60
		}
		if p[l] >= p[r] {
			p[l] += p[r]
			tp[ct] = x[r]
			ct--
			r--
		} else {
			p[r] += p[l]
			tp[ct] = x[l]
			ct--
			l++
		}
	}

	tp[1] = x[l]
	tp[0] = s
	ans := 0
	for i := 1; i <= n; i++ {
		if tp[i] > tp[i-1] {
			ans += tp[i] - tp[i-1]
		} else {
			ans += tp[i-1] - tp[i]
		}
	}
	fmt.Println(ans)
}
