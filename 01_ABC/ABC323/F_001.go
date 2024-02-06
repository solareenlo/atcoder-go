package main

import "fmt"

func check(a, b, c int) bool {
	if (b >= a && b >= c) || (b <= a && b <= c) {
		return true
	}
	return false
}

func main() {

	var xa, ya, xb, yb, xc, yc int
	fmt.Scan(&xa, &ya, &xb, &yb, &xc, &yc)
	ans := abs(xa-xb) + abs(ya-yb) - 1
	if check(xa, xb, xc) && check(ya, yb, yc) {
		ans += 2
		if xa == xb && xa == xc || ya == yb && ya == yc {
			ans += 2
		}
	}
	ans += abs(xb-xc) + abs(yb-yc)
	if xb != xc && yb != yc {
		ans += 2
	}
	if xb == xc && yb == yc {
		ans = 0
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
