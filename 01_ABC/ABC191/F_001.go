package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mp := map[int]int{}
	l := 1 << 30
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		for j := 1; j*j <= v; j++ {
			if v%j == 0 {
				mp[j] = gcd(mp[j], v)
				mp[v/j] = gcd(mp[v/j], v)
			}
		}
		l = min(l, v)
	}

	cnt := 0
	for k, v := range mp {
		if k == v && k <= l {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
