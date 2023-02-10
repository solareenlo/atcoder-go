package main

import "fmt"

const mod = 998244353

var ans int
var n, m, x, y, h int
var c []int
var s1, s2 [1505]int
var f []int

func main() {
	fmt.Scan(&n)
	n++

	ans = 1
	for x = 1; x < n; x++ {
		if (1 << lg2(x)) != x {
			calc()
		}
	}
	fmt.Println((ans + mod) % mod)
}

func calc() {
	y = (1 << (lg2(x) + 1)) - x
	if 2*x+y >= n {
		if n <= x+y {
			ans = (ans + 1) % mod
		} else {
			ans = (ans + n - x - y + 1) % mod
		}
		return
	}
	for h = 1; h < 12; h++ {
		if n < ((x + y) << (h - 1)) {
			break
		}
		dfs(1)
	}
}

func dfs(u int) {
	if u>>(h+1) != 0 {
		return
	}
	work(u)
	if u != 1 {
		dfs(u << 1)
	}
	dfs(u<<1 | 1)
}

func work(u int) {
	c = make([]int, 15)
	now := lg2(u)
	for i, j := u, 0; i>>1 > 0; i = j {
		j = i >> 1
		if i == (j<<1 | 1) {
			c[lg2(j)] = j - (1 << lg2(j)) + 1
		} else {
			c[lg2(j)] = j - (1 << lg2(j))
		}
	}
	c[now] = u - (1 << lg2(u)) + 1
	for j := u << 1; j>>(h+1) == 0; j = j<<1 | 1 {
		c[lg2(j)] = j - (1 << lg2(j)) + 1
	}
	m = n - x - 1
	for i := 0; i < h; i++ {
		m -= x * c[i]
	}
	m -= y * c[h]
	if now == h {
		m += y
	} else {
		m += x
	}
	if m < 0 {
		return
	}
	dp(now)
	if now < h {
		t := c[now]
		for i := 1; i <= x; i++ {
			tm := m - x - (t-1)*i
			if tm < 0 {
				break
			}
			v := tm + 2
			if i == x {
				if u == 1 {
					v--
				} else {
					ans = (ans + (x+1)*s1[tm]) % mod
					if tm >= i {
						ans = (ans - (x+1)*s1[tm-i]) % mod
					}
					continue
				}
			}
			ans = (ans + v*s1[tm] - s2[tm]) % mod
			if tm >= i {
				ans = (ans - v*s1[tm-i] + s2[tm-i]) % mod
			}
		}
		if u != 1 {
			for i := 0; i <= x; i++ {
				tm := m - (t-1)*i
				tl := x - i
				if tm < 0 {
					break
				}
				v := tm + 2
				ans = (ans + v*s1[tm] - s2[tm]) % mod
				if tm >= i {
					ans = (ans - v*s1[tm-i] + s2[tm-i]) % mod
					tm -= i
					ans = (ans + s1[tm]) % mod
					if tm >= tl {
						ans = (ans - s1[tm-tl]) % mod
					}
				}
			}
		}
	} else {
		ans = (ans + s1[m]) % mod
		if m >= y {
			ans = (ans - s1[m-y]) % mod
		}
		if u == (1<<(h+1))-1 {
			m -= y
			if m >= 0 {
				ans = (ans + (m+2)*s1[m] - s2[m]) % mod
				if m >= x {
					ans = (ans - (m+2)*s1[m-x] + s2[m-x]) % mod
				}
			}
		}
	}
}

func dp(now int) {
	f = make([]int, 1505)
	f[0] = 1
	for i := 0; i < h; i++ {
		if i == now {
			continue
		}
		v := c[i]
		for j := v; j <= m; j++ {
			f[j] = (f[j] + f[j-v]) % mod
		}
		v *= x + 1
		for j := m; j >= v; j-- {
			f[j] = (f[j] - f[j-v] + mod) % mod
		}
	}
	s1[0] = 1
	s2[0] = 0
	for i := 1; i <= m; i++ {
		s1[i] = (s1[i-1] + f[i]) % mod
		s2[i] = (s2[i-1] + i*f[i]) % mod
	}
}

func lg2(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}
