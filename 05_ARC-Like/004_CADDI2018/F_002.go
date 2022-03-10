package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	A := make([]int, N+1)
	for i := range A {
		A[i] = -1
	}

	const mod = 998244353
	pow2 := [100005]int{}
	pow2[0] = 1
	for i := 1; i <= N; i++ {
		pow2[i] = (pow2[i-1] * 2) % mod
	}

	Mp := make([]map[int]int, 100005)
	for i := range Mp {
		Mp[i] = map[int]int{}
	}
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if a == b {
			A[a] = c
			continue
		}
		if b < a {
			a, b = b, a
		}
		if _, ok := Mp[a][b]; ok {
			Mp[a][b] = ((c + Mp[a][b]) % 2) - 2
		} else {
			Mp[a][b] = c
		}
	}
	dp0 := 0
	dp1 := 0
	if A[N] == -1 || A[N] == 1 {
		dp1 = 1
	}
	if A[N] == -1 || A[N] == 0 {
		dp0 = 1
	}
	for i := N - 1; 1 <= i; i-- {
		Size := len(Mp[i])
		X := pow2[N-i-Size]
		x00 := X
		x01 := X
		x10 := X
		x11 := X
		for v, c := range Mp[i] {
			if 0 <= c {
				continue
			}
			if v == i+1 {
				if c != -2 {
					x00 = 0
					x11 = 0
				} else {
					x10 = 0
					x01 = 0
				}
			} else if v == i+2 {
				if c == -2 {
					x11 = 0
					x01 = 0
				} else {
					x00 = 0
					x10 = 0
				}
			} else {
				if c != -2 {
					x01 = 0
					x11 = 0
					x00 = 0
					x10 = 0
				}
			}
		}
		d1 := (x11*dp1 + x10*dp0) % mod
		d0 := (x01*dp1 + x00*dp0) % mod
		dp1 = d1
		dp0 = d0
		if A[i] == 1 {
			dp0 = 0
		}
		if A[i] == 0 {
			dp1 = 0
		}
	}
	fmt.Println((dp1 + dp0) % mod)
}
