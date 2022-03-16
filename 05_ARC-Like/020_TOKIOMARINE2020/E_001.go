package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var _n, k, S, T int
	fmt.Fscan(in, &_n, &k, &S, &T)
	if (S | T) != T {
		fmt.Println(0)
		return
	}

	const N = 55
	n := 0
	a := make([]int, N)
	for i := 0; i < _n; i++ {
		var x int
		fmt.Fscan(in, &x)
		if S == (S&x) && T == (T|x) {
			n++
			a[n] = x
		}
	}

	C := [N][N]int{}
	F := [N]int{}
	for i := 0; i <= n; i++ {
		C[i][0] = 1
		for j := 1; j <= i && j <= k; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
			F[i] += C[i][j]
		}
	}

	const M = 1 << 18
	c := [M]int{}
	SS := S ^ T
	ans := F[n]
	for S := SS; S > 0; S = (S - 1) & SS {
		res := 0
		for i := 1; i <= n; i++ {
			c[a[i]&S]++
		}
		for i := 1; i <= n; i++ {
			res += F[c[a[i]&S]]
			c[a[i]&S] = 0
		}
		if bits.OnesCount(uint(S))&1 != 0 {
			ans -= res
		} else {
			ans += res
		}
	}
	fmt.Println(ans)
}
