package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 1000005

type node struct{ x, v int }

var (
	n   int
	ans int
	t1  int = 1
	t2  int = 1
	ss  int
	k   int
	c   = make([]int, N)
	ci  = make([]int, N)
	h   = make([]int, N)
	a   = make([]node, 0)
)

func lowBit(x int) int {
	return x & -x
}

func query(x int) {
	for ; x <= n; x += lowBit(x) {
		ans = ans - 2*t2*c[x]
		ans %= mod
		ss += ci[x]
		ss %= mod
	}
}

func ins(x, v int) {
	y := h[x] * v
	x = a[x].x
	k += y
	k %= mod
	for ; x > 0; x -= lowBit(x) {
		c[x] += y
		c[x] %= mod
		if v > 0 {
			ci[x]++
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	a = make([]node, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].v)
		a[i].x = i
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].v < a[j].v
	})

	for i, j := 1, 1; i <= n; i++ {
		t := a[i].v - i + 1
		if t < 1 {
			fmt.Println(0)
			return
		}
		ans = (ans*t + t2*k) % mod
		query(a[i].x)
		t1 = t1 * t % mod
		h[i] = t1 * powMod(t2, mod-2) % mod
		if t > 1 {
			t2 = t2 * (t - 1) % mod
		}
		ins(i, 1)
		if t == 1 {
			for ; j <= i; j++ {
				ins(j, -1)
			}
		}
	}
	fmt.Println((((ans*powMod(2, mod-2)+ss*t1)%mod + mod) % mod))
}

const mod = 1000000007

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
