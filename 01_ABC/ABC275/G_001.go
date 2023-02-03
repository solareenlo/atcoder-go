package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const N = 200005

var n int
var a, b, c [N]float64

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
	}

	l, r := 0.0, 10.0
	ans := 0.0
	for i := 1; i <= 30; i++ {
		mid := (l + r) / 2
		if check(mid) {
			ans = mid
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Println(ans)
}

func check(mid float64) bool {
	mx1, mx2 := -1e18, -1e18
	for i := 1; i <= n; i++ {
		A := c[i] - mid*a[i]
		B := c[i] - mid*b[i]
		if A >= 0 && B >= 0 {
			return true
		}
		if A >= 0 {
			mx1 = math.Max(mx1, B/A)
		}
		if A <= 0 {
			mx2 = math.Max(mx2, -B/A)
		}
	}
	return mx1+mx2 >= 0
}
