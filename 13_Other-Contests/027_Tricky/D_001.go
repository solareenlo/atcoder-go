package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 1000000
	a := make([]int, N)
	b := make([]int, N)
	d := make([]bool, N)
	dp := make([]bool, N)
	for i := 0; i < 1000000; i++ {
		a[i] = -1
		b[i] = -1
	}

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var r int
		var n uint64
		fmt.Fscan(in, &n, &r)
		var tmp, s string
		fmt.Fscan(in, &tmp, &tmp, &tmp, &s, &tmp)
		x := uint64(0)
		for j := 0; ; j++ {
			if s[j] >= '0' && s[j] <= '9' {
				x *= 10
				x += uint64(s[j] - '0')
			} else {
				break
			}
		}
		mp := make(map[string]int)
		c := 0
		mp["Takahashi"] = c
		c++
		flag := false
		for j := 0; j < r; j++ {
			var t1, t2 string
			fmt.Fscan(in, &t1, &tmp, &t2)
			if flag {
				continue
			}
			if _, ok := mp[t1]; !ok {
				mp[t1] = c
				c++
			}
			if _, ok := mp[t2]; !ok {
				mp[t2] = c
				c++
			}
			p := mp[t1]
			q := mp[t2]
			if p == q {
				flag = true
				continue
			}
			if p == 0 {
				if x == 1 {
					flag = true
					continue
				}
				x--
			}
			if q == 0 {
				if x == n {
					flag = true
					continue
				}
				x++
			}
			if a[p] != -1 && a[p] != q {
				flag = true
				continue
			}
			if b[q] != -1 && b[q] != p {
				flag = true
				continue
			}
			if a[q] != -1 {
				b[a[q]] = p
			}
			if b[p] != -1 {
				a[b[p]] = q
			}
			a[p] = a[q]
			b[q] = b[p]
			a[q] = p
			b[p] = q
		}
		if uint64(c) > n {
			flag = true
		}
		if flag {
			fmt.Fprintln(out, "The memo must be wrong")
			for j := 0; j < c; j++ {
				a[j] = -1
				b[j] = -1
				d[j] = false
			}
			continue
		}
		y := 0
		m1 := x
		for {
			if d[y] {
				flag = true
				break
			}
			d[y] = true
			if a[y] == -1 {
				break
			}
			if m1 == 1 {
				flag = true
				break
			}
			m1--
			y = a[y]
		}
		d[0] = false
		y = 0
		m2 := x
		for {
			if d[y] {
				flag = true
				break
			}
			d[y] = true
			if b[y] == -1 {
				break
			}
			if m2 == n {
				flag = true
				break
			}
			m2++
			y = b[y]
		}
		if flag {
			fmt.Fprintln(out, "The memo must be wrong")
			for j := 0; j < c; j++ {
				a[j] = -1
				b[j] = -1
				d[j] = false
			}
			continue
		}
		v := make([]int, 0)
		sum := uint64(0)
		for j := 0; j < c; j++ {
			if !d[j] {
				x := j
				y := 0
				for {
					if d[x] {
						flag = true
						break
					}
					d[x] = true
					y++
					if a[x] == -1 {
						break
					}
					x = a[x]
				}
				d[j] = false
				x = j
				y--
				for {
					if d[x] {
						flag = true
						break
					}
					d[x] = true
					y++
					if b[x] == -1 {
						break
					}
					x = b[x]
				}
				if y > 1 {
					sum += uint64(y)
					v = append(v, y)
				}
			}
		}
		if flag {
			fmt.Fprintln(out, "The memo must be wrong")
			for j := 0; j < c; j++ {
				a[j] = -1
				b[j] = -1
				d[j] = false
			}
			continue
		}
		sort.Ints(v)
		m1--
		m2 = n - m2
		if m1 > m2 {
			m1, m2 = m2, m1
		}
		if sum > m2 {
			m := 0
			for j := uint64(0); j <= m1; j++ {
				dp[j] = false
			}
			dp[0] = true
			for j := 0; j < len(v); j++ {
				for k := min(m, int(m1)); k >= 0; k-- {
					if dp[k] {
						dp[k+v[j]] = true
						m = max(m, k+v[j])
					}
				}
			}
			var j uint64
			for j = sum - m2; j <= m1; j++ {
				if dp[j] {
					break
				}
			}
			if j > m1 {
				flag = true
			}
		}
		if flag {
			fmt.Fprintln(out, "The memo must be wrong")
		} else {
			if x%10 == 1 && x%100 != 11 {
				fmt.Fprintf(out, "Takahashi might get the %dst place\n", x)
			} else if x%10 == 2 && x%100 != 12 {
				fmt.Fprintf(out, "Takahashi might get the %dnd place\n", x)
			} else if x%10 == 3 && x%100 != 13 {
				fmt.Fprintf(out, "Takahashi might get the %drd place\n", x)
			} else {
				fmt.Fprintf(out, "Takahashi might get the %dth place\n", x)
			}
		}
		for j := 0; j < c; j++ {
			a[j] = -1
			b[j] = -1
			d[j] = false
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
