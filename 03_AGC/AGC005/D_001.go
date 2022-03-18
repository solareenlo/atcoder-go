package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	const skc = 924844033
	p := make([]int, 4004)
	p[0] = 1
	for i := 1; i <= n; i++ {
		p[i] = p[i-1] * i % skc
	}

	a := make([]int, 4004)
	a[0] = 1
	t := 0
	for i := 1; i <= n%m; i++ {
		t += n / m
		a[t] = 1
		t += n / m
		a[t] = 1
	}

	for i := 1; i <= m-n%m; i++ {
		t += n/m - 1
		a[t] = 1
		t += n/m - 1
		a[t] = 1
	}

	f := [4004][4004]int{}
	f[0][0] = 1
	for i := 1; i <= t; i++ {
		for j := 0; j <= n; j++ {
			if j > 0 {
				tmp := 0
				if a[i-1] == 0 {
					tmp = 1
				}
				f[i][j] = (f[i-1][j] + f[i-1-tmp][j-1]) % skc
			} else {
				f[i][j] = f[i-1][j] % skc
			}
		}
	}

	ans := 0
	for j := 0; j <= n; j++ {
		tmp := 1
		if j&1 != 0 {
			tmp = skc - 1
		}
		ans = (ans + f[t][j]*p[n-j]%skc*tmp) % skc
	}
	fmt.Println(ans)
}
