package main

import "fmt"

func main() {
	const MOD = 1000000007

	var h, w int
	fmt.Scan(&h, &w)
	if h > w {
		h, w = w, h
	}
	a := make([]int, 0)
	for i := 0; i < 1<<h; i++ {
		ok := true
		for j := 1; j <= i; j <<= 1 {
			if (i&j) != 0 && (i&(j<<1)) != 0 {
				ok = false
				break
			}
		}
		if ok {
			a = append(a, i)
		}
	}
	n := len(a)
	b := make([][]int, w+1)
	for i := range b {
		b[i] = make([]int, n)
	}
	b[0][0] = 1
	for i := 0; i < w; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				t := (a[j] & a[k]) | ((a[j] << 1) & a[k]) | (a[j] & (a[k] << 1))
				if t == 0 {
					b[i+1][k] = (b[i+1][k] + b[i][j]) % MOD
				}
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = (res + b[w][i]) % MOD
	}
	fmt.Println(res - 1)
}
