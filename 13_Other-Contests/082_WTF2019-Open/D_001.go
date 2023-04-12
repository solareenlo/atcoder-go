package main

import (
	"fmt"
	"math"
)

func main() {
	var R, B int
	fmt.Scan(&R, &B)
	if R > B {
		R, B = B, R
	}
	l := 3
	r := 2000000
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid, R, B) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	fmt.Println(l - 1)
}

type pair struct {
	x, y int
}

func check(k, X, Y int) bool {
	ans := pair{-1, -1}
	bns := ans
	tot := int(1e9 + 7)
	l := 0
	r := tot
	for l < r {
		mid := (l + r) >> 1
		now := Find(k, tot-mid, mid)
		if now.x <= X && now.y <= Y {
			return true
		}
		if now.x >= X && now.y >= Y {
			return false
		}
		if now.x < X {
			l = mid + 1
			ans = now
		} else {
			bns = now
			r = mid
		}
	}
	return ans.y*(bns.x-X)+bns.y*(X-ans.x) <= Y*(bns.x-ans.x)
}

func Find(k, x, y int) pair {
	r := int(math.Sqrt(2.0 * float64(k) * float64(x) * float64(y)))
	l := max(0, r-x-y-1)
	for l < r {
		mid := (l + r) >> 1
		cnt := Solve(mid/x+1, mid%x, x, y) + (mid/x + 1)
		if cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	X := 0
	Y := 0
	for i := 0; i*x <= r; i++ {
		c := (r-i*x)/y + 1
		X += c * i
		Y += c * (c - 1) / 2
	}
	return pair{X, Y}
}

func Solve(n, a, b, m int) int {
	if a >= m {
		return n*(a/m) + Solve(n, a%m, b, m)
	} else if b == 0 {
		return n * (a / m)
	} else if b >= m {
		return n*(n-1)/2*(b/m) + Solve(n, a, b%m, m)
	}
	return Solve((a+b*n)/m, (a+b*n)%m, m, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
