package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	c := make([]int, 100000)
	m := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if ((i > 0) && !(a > 0)) || !(i > 0) && (a > 0) {
			fmt.Println(0)
			return
		}
		m = max(m, a)
		c[a]++
	}

	ans := 1
	for i := 1; i <= m; i++ {
		ans *= powMod((powMod(2, c[i-1]) - 1), c[i])
		ans %= mod
		ans *= powMod(2, c[i]*(c[i]-1)/2)
		ans %= mod
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
