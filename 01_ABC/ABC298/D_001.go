package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)

	q := make([]int, 0)
	q = append(q, 1)
	val := 1
	for T > 0 {
		T--
		var op int
		fmt.Fscan(in, &op)
		if op == 1 {
			var x int
			fmt.Fscan(in, &x)
			val = (val*10 + x) % MOD
			q = append(q, x)
		} else if op == 2 {
			val = (val - q[0]*powMod(10, len(q)-1)%MOD + MOD) % MOD
			q = q[1:]
		} else {
			fmt.Fprintln(out, val)
		}
	}
}

const MOD = 998244353

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
