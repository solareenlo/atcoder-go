package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353
	const base = 10007

	var s string
	fmt.Fscan(in, &s)

	if len(s) == 1 {
		if s == "b" {
			fmt.Println(1, 0)
		} else {
			fmt.Println(2, 0)
		}
		return
	}

	var Pow [30000]int
	var F [50]int
	v := make([][]int, 50)

	hs := 0
	Pow[0] = 1
	for i := 0; i < len(s); i++ {
		hs = (hs*base + (int(s[i]-'a') + 1)) % MOD
		Pow[i+1] = Pow[i] * base % MOD
	}
	F[1] = 1
	F[2] = 1
	v[1] = append(v[1], 2)
	v[2] = append(v[2], 1)
	for i := 3; ; i++ {
		F[i] = F[i-1] + F[i-2]
		N := 1 << (i - 2)
		resize(&v[i], N)
		for j := 0; j < N; j++ {
			if j%2 == 0 {
				v[i][j] = (v[i-1][j/2]*Pow[F[i-2]] + v[i-2][j/4]) % MOD
			} else {
				v[i][j] = (v[i-2][j/4]*Pow[F[i-1]] + v[i-1][j/2]) % MOD
			}
			if len(s) == F[i] && v[i][j] == hs {
				fmt.Println(i, j)
				return
			}
		}
	}
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}
