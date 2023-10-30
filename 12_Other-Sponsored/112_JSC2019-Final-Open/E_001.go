package main

import (
	"bufio"
	"fmt"
	"os"
)

const MX = 500005

var a, b, cnt int
var val, q, fail [MX]int
var ch [MX][26]int

func insert(s string) {
	l := len(s) - 1
	p := 0
	val[0] = min(val[0], b*l)
	for i := 1; i <= l; i++ {
		if ch[p][s[i]-'a'] == 0 {
			cnt++
			ch[p][s[i]-'a'] = cnt
			val[cnt] = 1 << 60
		}
		p = ch[p][s[i]-'a']
		val[p] = min(val[p], b*l-a*i-b*i)
	}
}

func getnxt() {
	front, rear := 0, 0
	q[rear] = 0
	rear++
	for front < rear {
		t := q[front]
		front++
		val[t] = min(val[t], val[fail[t]])
		for i := 0; i < 26; i++ {
			if ch[t][i] != 0 {
				q[rear] = ch[t][i]
				rear++
				if t == 0 {
					fail[ch[t][i]] = 0
				} else {
					fail[ch[t][i]] = ch[fail[t]][i]
				}
			} else {
				if t == 0 {
					ch[t][i] = 0
				} else {
					ch[t][i] = ch[fail[t]][i]
				}
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	var s string

	val[0] = 1 << 60
	fmt.Fscan(in, &n, &m, &a, &b)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s)
		s = " " + s
		insert(s)
	}
	getnxt()
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &s)
		s = " " + s
		l := len(s) - 1
		p := 0
		ans := 1 << 60
		for j := 1; j <= l; j++ {
			p = ch[p][s[j]-'a']
			ans = min(ans, a*l+val[p])
		}
		fmt.Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
