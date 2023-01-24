package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353
const N = 200005
const HALF = (mod + 1) >> 1

var n int
var tr [N][3]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var m int
	fmt.Fscan(in, &n, &m)

	A := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &A[i])
		add(i, A[i])
	}

	var id, x, v int
	for m > 0 {
		fmt.Fscan(in, &id, &x)
		if id == 1 {
			fmt.Fscan(in, &v)
			add(x, v-A[x])
			A[x] = v
		} else {
			fmt.Fprintln(out, (sum(x, 2)*HALF%mod-(sum(x, 1)*(2*x+3)%mod)*HALF%mod+((sum(x, 0)*(x+1)%mod)*(x+2)%mod)*HALF%mod+mod*mod)%mod)
		}
		m--
	}
}

func add(x, c int) {
	for i := x; i <= n; i += i & -i {
		tr[i][0] = (tr[i][0] + c) % mod
		tr[i][1] = (tr[i][1] + (c * x % mod)) % mod
		tr[i][2] = (tr[i][2] + ((c * x % mod) * x % mod)) % mod
	}
}

func sum(x, t int) int {
	res := 0
	for i := x; i > 0; i -= i & -i {
		res = (res + tr[i][t]) % mod
	}
	return res
}
