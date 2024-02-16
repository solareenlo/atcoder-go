package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 20000

	var a, b [MX]int
	var used [MX]bool
	var v [MX][]int

	var n, K int
	fmt.Fscan(in, &n, &K)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
	}
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		cnt := 1
		x := i
		for {
			x = a[x]
			if x == i {
				break
			}
			cnt++
			used[x] = true
		}
		v[cnt] = append(v[cnt], i)
	}
	div := make([]int, 0)
	for i := 1; i*i <= K; i++ {
		if K%i == 0 {
			div = append(div, i)
			if i != K/i {
				div = append(div, K/i)
			}
		}
	}
	sort.Ints(div)
	for N := 1; N <= n; N++ {
		if len(v[N]) == 0 {
			continue
		}
		g := -1
		for _, i := range div {
			if gcd(N, K/i) == 1 {
				g = i
				break
			}
		}
		if len(v[N])%g != 0 {
			fmt.Println(0)
			return
		}
		for i := 0; i < len(v[N]); i += g {
			vs := make([]int, N*g)
			for j := i; j < i+g; j++ {
				x := v[N][j]
				p := j - i
				vs[p] = x
				for {
					x = a[x]
					p = (p + K) % (N * g)
					if x == v[N][j] {
						break
					}
					vs[p] = x
				}
			}
			for j := 0; j < len(vs); j++ {
				if j+1 < len(vs) {
					b[vs[j]] = vs[j+1]
				} else {
					b[vs[j]] = vs[0]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(b[i] + 1)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
