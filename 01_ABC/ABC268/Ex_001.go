package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 500005

var tr [N][26]int
var flg [N]int
var tot int

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	var m int
	fmt.Fscan(in, &s, &m)
	n := len(s)
	for i := 1; i <= m; i++ {
		var ch string
		fmt.Fscan(in, &ch)
		insert(ch)
	}

	q := make([]int, 0)
	for i := 0; i < 26; i++ {
		if tr[0][i] != 0 {
			q = append(q, tr[0][i])
		}
	}
	var fail [N]int
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		flg[u] |= flg[fail[u]]
		for i := 0; i < 26; i++ {
			if tr[u][i] != 0 {
				fail[tr[u][i]] = tr[fail[u]][i]
				q = append(q, tr[u][i])
			} else {
				tr[u][i] = tr[fail[u]][i]
			}
		}
	}

	ans := 0
	for i, p := 0, 0; i < n; i++ {
		c := int(s[i] - 'a')
		p = tr[p][c]
		if flg[p] != 0 {
			p = 0
			ans++
		}
	}
	fmt.Println(ans)
}

func insert(t string) {
	len := len(t)
	p := 0
	for i := 0; i < len; i++ {
		c := int(t[i] - 'a')
		if tr[p][c] == 0 {
			tot++
			tr[p][c] = tot
		}
		p = tr[p][c]
	}
	flg[p] = 1
}
