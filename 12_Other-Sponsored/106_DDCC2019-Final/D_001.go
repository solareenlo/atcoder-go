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

	var sum [5][5][1000100]int
	var disco string = "DISCO"

	MOD := pow(2, 32)

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			for k := 0; k < n; k++ {
				sum[i][j][k+1] = sum[i][j][k]
				if s[k] == disco[4-i+j] {
					if j != 0 {
						sum[i][j][k+1] = (sum[i][j][k+1] + sum[i][j-1][k]) % MOD
					} else {
						sum[i][j][k+1]++
					}
				}
			}
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		var ans [5]int
		for i := 0; i < 5; i++ {
			ans[i] = (ans[i] + (sum[i][i][r]-sum[i][i][l]+MOD)%MOD) % MOD
			for j := i + 1; j < 5; j++ {
				ans[j] = (ans[j] - sum[j][j-i-1][l]*ans[i]%MOD + MOD) % MOD
			}
		}
		fmt.Fprintln(out, ans[4])
	}
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}
