package main

import "fmt"

func main() {
	var n, k, r, s, p int
	var t string
	fmt.Scan(&n, &k, &r, &s, &p, &t)

	flag := make([]bool, 100001)
	res := 0
	for i := 0; i < n; i++ {
		if i < k || t[i-k] != t[i] || !flag[i-k] {
			if t[i] == 'r' {
				res += p
			} else if t[i] == 's' {
				res += r
			} else {
				res += s
			}
			flag[i] = true
		}
	}
	fmt.Println(res)
}
