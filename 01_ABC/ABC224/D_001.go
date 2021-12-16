package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &m)

	g := make([][]int, 10)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	ss := strings.Split("999999999", "")
	for i := 1; i < 9; i++ {
		var p int
		fmt.Fscan(in, &p)
		ss[p-1] = string('0' + i)
	}
	s := strings.Join(ss, "")

	q := make([]string, 0)
	q = append(q, s)
	mp := map[string]int{}
	mp[s] = 0

	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		v := 0
		for i := 1; i < 10; i++ {
			if s[i-1] == '9' {
				v = i
			}
		}
		for _, u := range g[v] {
			tt := strings.Split(s, "")
			tt[u-1], tt[v-1] = tt[v-1], tt[u-1]
			t := strings.Join(tt, "")
			if _, ok := mp[t]; ok {
				continue
			}
			mp[t] = mp[s] + 1
			q = append(q, t)
		}
	}

	if i, ok := mp["123456789"]; ok {
		fmt.Println(i)
	} else {
		fmt.Println(-1)
	}
}
