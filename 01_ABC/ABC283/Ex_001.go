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

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n, m, r int
		fmt.Fscan(in, &n, &m, &r)
		l := (n + m - r) / m
		ans := l*(l-1)/2*m + l*r
		for d := 1; d < 30; d++ {
			ans -= floor_sum(l, 1<<d, m, r)
		}
		fmt.Fprintln(out, ans)
	}
}

func floor_sum(n, m, a, b int) int {
	ans := 0
	if a >= m {
		ans += (n - 1) * n * (a / m) / 2
		a %= m
	}
	if b >= m {
		ans += n * (b / m)
		b %= m
	}

	yMax := (a*n + b) / m
	xMax := (yMax*m - b)
	if yMax == 0 {
		return ans
	}
	ans += (n - (xMax+a-1)/a) * yMax
	ans += floor_sum(yMax, a, m, (a-xMax%a)%a)
	return ans
}
