package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 8080

	var n int
	fmt.Fscan(in, &n)
	g := make([]int, N)
	f := make([]int, N)
	for i := 0; i <= n; i++ {
		g[i] = -i * (i + 1) / 2
		f[i] = -i * (i + 1) / 2
	}

	var s, m, c1, c2 int
	v := make([]int, N)
	sum := 0
	for i := 1; i <= n; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		if a != c && b != d {
			if (a ^ b) != 0 {
				s++
				sum += i
			} else {
				m++
				v[m] = i
			}
		} else if a == c && b == d {
			if a == 2 {
				sum += i
				c1++
			}
			if b == 2 {
				sum += i
				c2++
			}
		} else if a == c {
			if a == 2 {
				sum += i
				c1++
			}
			for j := 0; j < n; j++ {
				g[j] = max(g[j], g[j+1]+i)
			}
		} else {
			if b == 2 {
				sum += i
				c2++
			}
			for j := 0; j < n; j++ {
				f[j] = max(f[j], f[j+1]+i)
			}
		}
	}
	tmp := v[1 : m+1]
	tmp = reverseOrderInt(tmp)
	for i := 0; i < m; i++ {
		v[i+1] = tmp[i]
	}

	ans := 0
	for i := 0; i <= m; i++ {
		sum += 2 * v[i]
		for j := 0; j <= s; j++ {
			ans = max(ans, sum+f[c1+j]+g[c2+s-j])
		}
		c1++
		c2++
	}
	fmt.Println(ans)
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
