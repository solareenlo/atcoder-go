package main

import "fmt"

func main() {
	const N = 200200
	a := [N][33]int{}

	var s string
	fmt.Scan(&s)
	n := len(s)
	for i := 0; i < 26; i++ {
		a[n][i] = n + 1
	}

	f := make([]int, N)
	p := make([]int, N)
	f[n] = 1
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			a[i][j] = a[i+1][j]
		}
		a[i][s[i]-97] = i + 1
		f[i] = N
		for j := 0; j < 26; j++ {
			if f[i] > f[a[i][j]]+1 {
				f[i] = f[a[i][j]] + 1
				p[i] = j
			}
		}
	}

	for i := 0; i < n; i = a[i][p[i]] {
		fmt.Print(string(p[i] + 97))
	}
}
