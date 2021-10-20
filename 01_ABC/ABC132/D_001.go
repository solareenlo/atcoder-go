package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	mod := int(1e9 + 7)
	c := [2001][2001]int{}
	c[0][0] = 1
	for i := 1; i < n+1; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 {
				c[i][j] = 1
			} else {
				c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
			}
		}
	}

	for i := 0; i < k; i++ {
		fmt.Println(c[n-k+1][i+1] * c[k-1][i] % mod)
	}
}
