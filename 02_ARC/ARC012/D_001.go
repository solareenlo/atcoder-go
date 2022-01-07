package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sum = [100010]int{}
	mod int
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscan(in, &n, &t, &mod)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x = abs(x)
		y = abs(y)
		if (t+x+y)%2 != 0 || t-x-y < 0 {
			fmt.Println(0)
			return
		}
		m := (t - x - y) / 2
		add(t, 2)
		add(m, -1)
		add(t-m, -1)
		add(x+m, -1)
		add(t-x-m, -1)
	}
	for i := 1; i <= t; i++ {
		sum[i] += sum[i-1]
	}
	for i := t; i >= 2; i-- {
		for j := 2 * i; j <= t; j += i {
			sum[i] += sum[j]
			sum[j/i] += sum[j]
			sum[j] = 0
		}
	}

	ans := 1
	for i := 2; i <= t; i++ {
		ans *= powMod(i, sum[i])
		ans %= mod
	}
	fmt.Println(ans)
}

func add(i, a int) {
	sum[1] += a
	sum[i+1] -= a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

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
