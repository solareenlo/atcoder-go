package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, c, k int
	fmt.Fscan(in, &n, &c, &k)
	g := make([][]int, 100)
	for i := 0; i < n; i++ {
		var d, v int
		fmt.Fscan(in, &d, &v)
		g[d] = append(g[d], v)
	}

	s1 := 1
	var a [7001000]int
	var fz [7100000]int
	const G = 100000
	a[0] = 0
	ff := c + 1
	for i := 1; i <= c; i++ {
		p := 0
		for _, j := range g[i] {
			for o := 0; o < s1; o++ {
				fz[p] = a[o] ^ j
				p++
			}
		}
		s1 = p
		for j := 0; j < s1; j++ {
			a[j] = fz[j]
		}
		if s1 > G {
			ff = i + 1
			break
		}
	}

	s2 := 1
	var b [7010010]int
	b[0] = 0
	for i := ff; i <= c; i++ {
		p := 0
		for _, j := range g[i] {
			for o := 0; o < s2; o++ {
				fz[p] = b[o] ^ j
				p++
			}
		}
		s2 = p
		for j := 0; j < s2; j++ {
			b[j] = fz[j]
		}
	}

	var ch [20100000][2]int
	var su [20100000]int
	cn := 1
	for i := 0; i < s1; i++ {
		u := 1
		for j := 60; j >= 0; j-- {
			if ch[u][(a[i]>>j)&1] == 0 {
				cn++
				ch[u][(a[i]>>j)&1] = cn
			}
			u = ch[u][(a[i]>>j)&1]
			su[u]++
		}
	}

	var ps [7010000]int
	ans := 0
	for i := 0; i < s2; i++ {
		ps[i] = 1
	}
	for i := 60; i >= 0; i-- {
		sum := 0
		for j := 0; j < s2; j++ {
			tmp := 0
			if (b[j]>>i)&1 == 0 {
				tmp = 1
			}
			sum += su[ch[ps[j]][tmp]]
		}
		if k > sum {
			k -= sum
			for j := 0; j < s2; j++ {
				ps[j] = ch[ps[j]][(b[j]>>i)&1]
			}
		} else {
			ans |= (1 << i)
			for j := 0; j < s2; j++ {
				tmp := 0
				if (b[j]>>i)&1 == 0 {
					tmp = 1
				}
				ps[j] = ch[ps[j]][tmp]
			}
		}
	}

	fmt.Println(ans)
}
