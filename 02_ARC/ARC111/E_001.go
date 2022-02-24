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

	for i := 0; i < t; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		e := (d + c - b - 2) / (c - b)
		ans := floor_sum(e, d, c, a) - floor_sum(e, d, b, a-1)
		fmt.Fprintln(out, e-1-ans)
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
