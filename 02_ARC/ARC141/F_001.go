package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 2000005
const P = 13331

type node struct {
	nxt             []int
	c, l, fail, len int
	v               int
	ok              bool
}

var s []string
var t []node
var f, h, Q []int
var tot int
var v, v2 []string

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	s = make([]string, 1000005)
	f = make([]int, N)
	h = make([]int, N)
	t = make([]node, N)
	for i := range t {
		t[i].nxt = make([]int, 4)
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		extend(s[i])
	}
	Q = make([]int, 0)
	build()
	v = make([]string, 0)
	v2 = make([]string, 0)
	for i := 1; i <= n; i++ {
		if check(s[i]) {
			v = append(v, s[i])
		} else {
			v2 = append(v2, s[i])
		}
	}
	t = make([]node, N)
	for i := range t {
		t[i].nxt = make([]int, 4)
	}
	tot = 0
	for i := 0; i < len(v); i++ {
		extend(v[i])
	}
	build()
	for i := 0; i < len(v); i++ {
		check2(v[i])
	}
	for i := 0; i < len(v2); i++ {
		check3(v2[i])
	}
	fmt.Println("No")
}

var stk []int

func check3(s string) {
	stk = make([]int, 0)
	stk = append(stk, 0)
	for i := 0; i < len(s); i++ {
		it := stk[len(stk)-1]
		it = t[it].nxt[s[i]-'A']
		stk = append(stk, it)
		if t[it].c != 0 {
			for i := 1; i <= t[it].l; i++ {
				if len(stk) > 0 {
					stk = stk[:len(stk)-1]
				}
			}
		}
	}
	if len(stk) != 1 {
		fmt.Println("Yes")
		os.Exit(0)
	}
}

func check2(s string) {
	it := 0
	l := len(s)
	f[0] = 1
	for i := 1; i <= l; i++ {
		f[i] = f[i-1] * P
		h[i] = h[i-1]*P + int(s[i-1])
	}
	for i := 0; i < l; i++ {
		it = t[it].nxt[s[i]-'A']
	}
	for it = t[it].fail; it > 0; it = t[it].fail {
		if t[it].ok && t[it].v != ask(1, l-t[it].len) {
			fmt.Println("Yes")
			os.Exit(0)
		}
	}
}

func check(s string) bool {
	it := 0
	l := len(s)
	for i := 0; i < l; i++ {
		it = t[it].nxt[s[i]-'A']
		if t[it].c > 1 || (t[it].c > 0 && i+1 != l) {
			return false
		}
	}
	return true
}

func build() {
	for i := 0; i < 4; i++ {
		if t[0].nxt[i] != 0 {
			Q = append(Q, t[0].nxt[i])
		}
	}
	for len(Q) > 0 {
		x := Q[0]
		y := t[x].fail
		Q = Q[1:]
		t[x].c += t[y].c
		if t[y].c != 0 {
			t[x].l = t[y].l
		}
		for i := 0; i < 4; i++ {
			if t[x].nxt[i] != 0 {
				Q = append(Q, t[x].nxt[i])
				t[t[x].nxt[i]].fail = t[y].nxt[i]
			} else {
				t[x].nxt[i] = t[y].nxt[i]
			}
		}
	}
}

func extend(s string) {
	l := len(s)
	f[0] = 1
	for i := 1; i <= l; i++ {
		f[i] = f[i-1] * P
		h[i] = h[i-1]*P + int(s[i-1])
	}
	it := 0
	for i := 0; i < l; i++ {
		if t[it].nxt[s[i]-'A'] == 0 {
			tot++
			t[it].nxt[s[i]-'A'] = tot
			t[tot].len = i + 1
		}
		it = t[it].nxt[s[i]-'A']
		if i+1 != l {
			x := ask(i+2, l)
			if !t[it].ok {
				t[it].ok = true
				t[it].v = x
			} else if t[it].v != x {
				t[it].v = 0
			}
		}
	}
	t[it].c++
	t[it].l = l
}

func ask(l, r int) int {
	return h[r] - h[l-1]*f[r-l+1]
}
