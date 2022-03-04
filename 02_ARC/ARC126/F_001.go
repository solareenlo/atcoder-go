package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 998244353
const inv2 = (mod + 1) >> 1

func add(a, b int) int {
	if a+b >= mod {
		return a + b - mod
	}
	return a + b
}
func mul(a, b int) int { return a * b % mod }

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

type num struct{ p, q, id int }

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	x := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}
	x[n+1] = x[1]

	d := make([]int, n+1)
	sum := 0
	tim := 0
	m := 0
	a := make([]num, 1000005)
	for i := 1; i <= n; i++ {
		d[i] = x[i+1] - x[i]
		if d[i] < 0 {
			sum = add(sum, 1)
			tmp := 0
			if n == i {
				tmp = 1
			}
			tim = add(tim, tmp)
		}
		val := abs(d[i])
		for j := 1; j < val; j++ {
			m++
			a[m].p = j
			a[m].q = val
			a[m].id = i
		}
	}
	m++
	a[m].p = 1
	a[m].q = 1
	tmp := a[1 : m+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].p*tmp[j].q < tmp[j].p*tmp[i].q
	})

	ans := 0
	for i, pre := 1, 0; i <= m; i++ {
		now := mul(a[i].p, powMod(a[i].q, mod-2))
		if sum == 1 {
			res := mul(inv2, mul(add(mul(now, now), mod-mul(pre, pre)), add(d[n], mod)))
			res = add(res, mul(add(now, mod-pre), tim))
			ans = add(ans, res)
		}
		pre = now
		if d[a[i].id] > 0 {
			tmp := 0
			if n == a[i].id {
				tmp = 1
			}
			sum = add(sum, mod-1)
			tim = add(tim, mod-tmp)
		} else {
			tmp := 0
			if n == a[i].id {
				tmp = 1
			}
			sum = add(sum, 1)
			tim = add(tim, tmp)
		}
	}
	fmt.Println(mul(ans, powMod(3, mod-2)))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
