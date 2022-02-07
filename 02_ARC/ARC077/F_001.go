package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	k   int
	w   = make([]int, 200200)
	fib = make([]int, 200200)
)

func dfs(r, v int) {
	if r <= n {
		w[r] += v
		return
	}
	p := 0
	for ; fib[p+1]*n+fib[p]*k < r; p++ {

	}
	dfs(r-fib[p]*n-fib[p-1]*k, v)
	w[k] += v * fib[p-1]
	w[n] += v * fib[p]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)
	n = len(s) / 2
	s = " " + s

	var l, r int
	fmt.Fscan(in, &l, &r)

	fib[1] = 1
	for i := 2; fib[i-1] <= int(2e18); i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	p := 0
	nxt := make([]int, 200200)
	for i := 2; i <= n; i++ {
		for p != 0 && s[p+1] != s[i] {
			p = nxt[p]
		}
		if s[p+1] == s[i] {
			p++
		}
		nxt[i] = p
	}

	ans := [26]int{}
	k = n - p
	dfs(l-1, -1)
	dfs(r, 1)
	for i := n; i >= 1; i-- {
		w[i-1] += w[i]
		ans[s[i]-'a'] += w[i]
	}

	for i := 0; i < 26; i++ {
		fmt.Print(ans[i], " ")
	}
	fmt.Println()
}
