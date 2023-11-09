package main

import "fmt"

func cnt(b, d, k int) int {
	if d == 0 {
		if b >= 0 {
			return k
		}
		return 0
	} else if d > 0 {
		if b >= 0 {
			return k
		} else {
			return max(k-((-b+d-1)/d), 0)
		}
	}
	if b < 0 {
		return 0
	}
	return min(k, b/(-d)+1)
}

func main() {
	const p = 998244353
	const N = 200010

	var pre [N]int

	var n, k int
	fmt.Scan(&n, &k)
	var s string
	fmt.Scan(&s)
	s = " " + s
	for i := 1; i <= n; i++ {
		tmp := 0
		if s[i] == 'R' {
			tmp = 1
		}
		pre[i] = pre[i-1] + (1 - tmp*2)
	}
	tot := pre[n]
	if tot == 0 {
		ans := 0
		for i := 1; i <= n; i++ {
			if pre[i-1] < 0 || pre[i] < 0 {
				ans++
			}
		}
		fmt.Println(ans * k % p)
		return
	} else if tot > 0 {
		totL := 0
		totR := 0
		for i := 1; i <= n; i++ {
			if s[i] == 'L' {
				totL += cnt(-pre[i], -tot, k)
			} else {
				totR += k
			}
		}
		pos := n
		for pre[pos] >= pre[n] {
			pos--
		}
		totR -= (n - pos) / 2
		fmt.Println((totL*2 + totR) % p)
	} else {
		totL := 0
		totR := 0
		for i := 1; i <= n; i++ {
			if s[i] == 'R' {
				totR += cnt(pre[n]-pre[i-1], tot, k)
			} else {
				totL += k
			}
		}
		pos := 0
		for pre[pos] >= 0 {
			pos++
		}
		totL -= pos / 2
		fmt.Println((totR*2 + totL) % p)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
