package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var g [1 << 17][]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[v] = append(g[v], u)
	}
	var s string
	fmt.Fscan(in, &s)
	ni := 0
	for s[ni] == '"' {
		ni++
	}
	ni += 5
	now := 0
	for ni < len(s) && '0' <= s[ni] && s[ni] <= '9' {
		now *= 10
		now += int(s[ni] - '0')
		ni++
	}
	now--
	cnt := make([]int, n)
	for _, to := range g[now] {
		cnt[to]++
	}
	for i := 0; i < n; i++ {
		cnt[i] ^= 1
	}
	if ni < len(s) && s[ni] == 'w' {
		ni++
		for i := 0; i < n; i++ {
			cnt[i] ^= 1
		}
	}
	for ni < len(s) {
		mask := 1
		ni++
		nxt := make([]int, n)
		if ni < len(s) && s[ni] == 'w' {
			ni += 2
			mask = 0
		}
		aa := 0
		for i := 0; i < n; i++ {
			if cnt[i] == 1 {
				aa++
				for _, to := range g[i] {
					nxt[to]++
				}
			}
		}
		for i := 0; i < n; i++ {
			if mask == 0 {
				if nxt[i] > 0 {
					cnt[i] = 1
				} else {
					cnt[i] = 0
				}
			} else {
				if nxt[i] != aa {
					cnt[i] = 1
				} else {
					cnt[i] = 0
				}
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += cnt[i]
	}
	fmt.Println(ans)
}
