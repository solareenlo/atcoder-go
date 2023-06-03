package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var n, q int
	fmt.Fscan(in, &n, &q)
	var s string
	fmt.Fscan(in, &s)

	v := make([]int, n+1)
	w := make([]int, n+1)
	w[0] = 1
	for i := 0; i < n; i++ {
		v[i+1] = (v[i]*100 + (int(s[i]-'a') + 1)) % MOD
		w[i+1] = (w[i] * 100) % MOD
	}

	for q > 0 {
		q--
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		val1 := (v[b] - v[a-1]*w[b-a+1]%MOD + MOD) % MOD
		val2 := (v[d] - v[c-1]*w[d-c+1]%MOD + MOD) % MOD
		if val1 == val2 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
