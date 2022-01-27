package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, c, k int
	fmt.Fscan(in, &n, &c, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	if c == 1 {
		for i := 0; i < n; i++ {
			fmt.Fprintln(out, a[i])
		}
		return
	}

	for k > 0 && c*a[0] <= a[n-1] {
		a[0] *= c
		k--
		sort.Ints(a)
	}

	for i := 0; i < n; i++ {
		a[i] = a[i] % MOD * powMod(c, k/n) % MOD
	}
	for i := 0; i < k%n; i++ {
		a[i] = a[i] * c % MOD
	}
	a = append(a[k%n:], a[:k%n]...)

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, a[i])
	}
}

const MOD = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}
