package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var a, b int
var fact, ifact, inv [200111]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &a, &b, &n)
	if a == b {
		fmt.Println(0)
		return
	}
	pts := make([]pair, 0)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if (x == a && y == b) || (x == 0 && y == 0) {
			fmt.Println(0)
			return
		}
		pts = append(pts, pair{x, y})
		pts = append(pts, pair{a - x, b - y})
	}
	fact[0] = 1
	ifact[0] = 1
	for i := 1; i <= 200011; i++ {
		inv[i] = modpow(i, MOD-2)
		fact[i] = mult(fact[i-1], i)
		ifact[i] = mult(ifact[i-1], inv[i])
	}
	pts = append(pts, pair{a, b})
	pts = append(pts, pair{0, 0})
	pts = uniquePair(pts)
	n = len(pts)
	var dp [51]int
	dp[0] = 1
	for i := 1; i < n; i++ {
		ans := ways(a, b, pts[i].x, pts[i].y)
		for j := 1; j < i; j++ {
			ans = subtr(ans, mult(dp[j], ways(pts[j].x, pts[j].y, pts[i].x, pts[i].y)))
		}
		dp[i] = ans
	}
	fmt.Println(dp[n-1])
}

type pair struct {
	x, y int
}

func uniquePair(a []pair) []pair {
	occurred := map[pair]bool{}
	result := []pair{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].x == result[j].x {
			return result[i].y > result[j].y
		}
		return result[i].x > result[j].x
	})
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func choose(a, b int) int {
	if a < b {
		return 0
	}
	if b < 0 {
		return 0
	}
	if b == 0 {
		return 1
	}
	if a == b {
		return 1
	}
	return mult(fact[a], mult(ifact[b], ifact[a-b]))
}

func ways(x1, y1, x2, y2 int) int {
	if x2 > x1 || y2 > y1 {
		return 0
	}
	if x1 < y1 {
		return 0
	}
	if x2 < y2 {
		return 0
	}
	if x1-y1 > a-b {
		return 0
	}
	if x2-y2 > a-b {
		return 0
	}
	if x1 == x2 && y1 == y2 {
		return 1
	}
	d1 := y1 - y2
	d2 := x1 - x2
	ans := 0
	t := d1 + d2
	i := d1 - d2
	A := 1 + (x2 - y2)
	B := (a - b + 1) - (x2 - y2)
	for l := -(((t+i)/2 + B) / (A + B)) - 1; l <= (t-((t+i)/2))/(A+B)+1; l++ {
		ans = add(ans, subtr(choose(t, (t+i)/2+l*(A+B)), choose(t, (t+i)/2+B+l*(A+B))))
	}
	return ans
}

const MOD = 1000000007

func add(a, b int) int {
	a += b
	for a >= MOD {
		a -= MOD
	}
	return a
}

func subtr(a, b int) int {
	return add(a, MOD-b)
}

func mult(a, b int) int {
	return a * b % MOD
}

func modpow(a, b int) int {
	r := 1
	for b != 0 {
		if (b & 1) != 0 {
			r = mult(r, a)
		}
		a = mult(a, a)
		b >>= 1
	}
	return r
}
