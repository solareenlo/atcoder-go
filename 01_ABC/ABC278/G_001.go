package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 2020

var l int
var SG [N]int
var v [N]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, r int
	fmt.Fscan(in, &n, &l, &r)

	var x, y, tmp int
	if (l^r) != 0 || (((n ^ l) & 1) == 0) {
		fmt.Fprintln(out, "First")
		out.Flush()
		tmp = l + ((n - l) & 1)
		fmt.Fprintln(out, (n-tmp)/2+1, tmp)
		out.Flush()
		fmt.Fscan(in, &x, &y)
		n -= tmp
		for x > 0 && y > 0 {
			if x <= n/2 {
				fmt.Fprintln(out, x+n/2+tmp, y)
				out.Flush()
			} else {
				fmt.Fprintln(out, x-n/2-tmp, y)
				out.Flush()
			}
			fmt.Fscan(in, &x, &y)
		}
	} else {
		for i := 0; i <= n; i++ {
			SG[i] = -1
			v[i] = true
		}
		dfs(n)
		if SG[n] != 0 {
			fmt.Fprintln(out, "First")
			out.Flush()
			x, y = 1, 1
		} else {
			fmt.Fprintln(out, "Second")
			out.Flush()
			fmt.Fscan(in, &x, &y)
			op(x, y)
		}
		for x > 0 && y > 0 {
			tmp = 0
			for i, c := 1, 0; i <= n; i++ {
				if v[i] {
					c = i
					for v[c] {
						c++
					}
					c--
					tmp ^= dfs(c - i + 1)
					i = c
				}
			}
			for i, c := 1, 0; i <= n && tmp != 0; i++ {
				if v[i] {
					c = i
					for v[c] {
						c++
					}
					c--
					for di := 0; di+l <= c-i+1; di++ {
						if (dfs(c-i+1) ^ dfs(di) ^ dfs(c-i+1-l-di)) == tmp {
							fmt.Fprintln(out, i+di, l)
							out.Flush()
							op(i+di, l)
							tmp = 0
							break
						}
					}
					i = c
				}
			}
			fmt.Fscan(in, &x, &y)
			op(x, y)
		}
		return
	}
}

func dfs(c int) int {
	if SG[c] != -1 {
		return SG[c]
	}
	if c < l {
		SG[c] = 0
		return SG[c]
	}
	var cnt [N]int
	for i := 0; i*2 <= c-l; i++ {
		cnt[dfs(i)^dfs(c-l-i)] = 1
	}
	SG[c] = 0
	for cnt[SG[c]] != 0 {
		SG[c]++
	}
	return SG[c]
}

func op(st, c int) {
	for i := st; i <= st+c-1; i++ {
		v[i] = false
	}
}
