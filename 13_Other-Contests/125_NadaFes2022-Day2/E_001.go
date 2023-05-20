package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 233333
const MOD = 998244353

var n, ans int
var fac [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fac[0] = 1
	for i := 1; i < N; i++ {
		fac[i] = fac[i-1] * i % MOD
	}

	fmt.Fscan(in, &n)

	x := make([]int, n+1)
	y := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	if n == 2 {
		fmt.Println(2 * min(abs(x[1]-x[2]), abs(y[1]-y[2])))
		return
	}

	ans = (ans + calc(x)*2 + calc(y)*2) % MOD
	for i := 1; i < n+1; i++ {
		a := x[i]
		b := y[i]
		x[i] = a + b
		y[i] = a - b
	}

	ans = (ans - calc(x) - calc(y) + 2*MOD) % MOD
	fmt.Println(ans)
}

func calc(b []int) int {
	a := make([]int, len(b))
	copy(a, b)
	tmp := a[1:]
	sort.Ints(tmp)
	res := 0
	for i := 3; i <= n; i++ {
		res = (res + (i-1)*(i-2)%MOD*fac[n-3]%MOD*a[i]%MOD + MOD) % MOD
	}
	for i := 1; i < n-1; i++ {
		res = (res - (n-i)*(n-i-1)%MOD*fac[n-3]%MOD*a[i]%MOD + MOD) % MOD
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
