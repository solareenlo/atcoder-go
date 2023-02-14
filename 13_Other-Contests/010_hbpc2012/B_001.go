package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	nxt, to, w int
}

const mod = 1440

var in = bufio.NewReader(os.Stdin)
var n, tot, cnt int
var s1, s2 string
var mp map[string]int
var ans []string
var e []edge
var head []int
var ck []bool

func main() {
	const N = 2020

	fmt.Fscan(in, &s1)
	tot++
	mp = make(map[string]int)
	mp[s1] = tot
	ans = make([]string, 27)
	for i := 0; i <= 23; i++ {
		fmt.Fscan(in, &ans[i])
	}
	e = make([]edge, N)
	head = make([]int, N)
	ck = make([]bool, N)
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s1)
		t1 := read()
		fmt.Fscan(in, &s2)
		t2 := read()
		if mp[s1] == 0 {
			tot++
			mp[s1] = tot
		}
		if mp[s2] == 0 {
			tot++
			mp[s2] = tot
		}
		add(mp[s1], mp[s2], (t2-t1+mod)%mod)
		add(mp[s2], mp[s1], (t1-t2+mod)%mod)
	}
	fmt.Fscan(in, &s1)
	var h, m int
	fmt.Fscanf(in, "\n%02d:%02d", &h, &m)
	dfs(mp[s1], h*60+m)
}

func dfs(u, res int) {
	if ck[u] {
		return
	}
	ck[u] = true
	if u == 1 {
		fmt.Println(ans[res/60])
	}
	for i := head[u]; i > 0; i = e[i].nxt {
		v := e[i].to
		dfs(v, (res+e[i].w)%mod)
	}
}

func read() int {
	var h, m int
	fmt.Fscanf(in, "%02d:%02d", &h, &m)
	return h*60 + m
}

func add(u, v, ww int) {
	cnt++
	e[cnt].nxt = head[u]
	e[cnt].to = v
	e[cnt].w = ww
	head[u] = cnt
}
