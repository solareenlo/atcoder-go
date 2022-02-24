package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n int
)

var in = bufio.NewReader(os.Stdin)

func cal() int {
	sum := 0
	p := 1
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		sum += x * p
		p *= 2 * (i + 1)
	}
	return sum
}

func main() {
	var m int
	fmt.Fscan(in, &n, &m)

	x := cal()
	a := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = cal()
	}

	t := 1
	for i := 1; i <= n; i++ {
		t *= 2 * i
	}
	t--

	g := t
	for i := 0; i < m; i++ {
		g = gcd(g, a[i])
	}

	e := make([]int, 55)
	if g < t/g {
		e[0] = 1
		for i := 0; i < n; i++ {
			e[i+1] = e[i] * 2 * (i + 1)
		}
		d := make([]int, g)
		q := make([]int, 0)
		for i := 0; i <= n; i++ {
			d[e[i]%g] = 1
			q = append(q, e[i]%g)
		}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for i := 0; i <= n; i++ {
				v := (u + e[i]) % g
				if d[v] == 0 {
					d[v] = d[u] + 1
					q = append(q, v)
				}
			}
		}
		fmt.Println(d[x%g])
	} else {
		ans := 1 << 60
		for i := x % g; i <= t; i += g {
			if i == 0 {
				continue
			}
			r := i
			sum := 0
			for j := 1; j <= n; j++ {
				sum += r % (2 * j)
				r /= (2 * j)
			}
			ans = min(ans, sum)
		}
		fmt.Println(ans)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
