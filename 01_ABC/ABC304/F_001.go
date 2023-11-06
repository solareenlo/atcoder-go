package main

import "fmt"

func main() {
	const P = 998244353
	const N = 200200

	var v [N]int

	var n int
	var s string
	fmt.Scan(&n, &s)
	s = " " + s
	ans := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			c := 1
			for j := 1; j <= i; j++ {
				f := true
				for k := j; k <= n; k += i {
					if s[k] == '.' {
						f = false
						break
					}
				}
				if f {
					c = (c * 2) % P
				}
			}
			ans = (ans + c - v[i] + P) % P
			for j := i + i; j <= n; j += i {
				v[j] = (v[j] + c - v[i] + P) % P
			}
		}
	}
	fmt.Println(ans)
}
