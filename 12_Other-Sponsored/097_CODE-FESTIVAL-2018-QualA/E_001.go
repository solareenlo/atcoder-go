package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAXN = 200020

type oranges struct {
	val, id, cnt int
}

var v []oranges
var seen, mx, mn [MAXN]int
var bad, LB, UB, X int

func add(id, cnt int) {
	if seen[id] == -1 {
		bad--
		mn[id] = cnt
		mx[id] = cnt
		if cnt == 0 {
			seen[id] = 1
		} else {
			seen[id] = 2
		}
		LB += cnt
		UB += cnt
		return
	}
	if mn[id] > mx[id] {
		bad--
	}
	if seen[id] == 1 {
		UB++
		mx[id]++
	} else {
		LB--
		mn[id]--
	}
}

func del(id, cnt int) {
	if seen[id] == 1 {
		LB++
		mn[id]++
	} else {
		UB--
		mx[id]--
	}
	if mn[id] > mx[id] {
		bad++
	}
}

func ok() bool {
	if bad != 0 {
		return false
	}
	return LB <= X && X <= UB
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, B [MAXN]int

	var Y, n int
	fmt.Fscan(in, &X, &Y, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}
	each := (X + Y) / n
	for i := 1; i <= n; i++ {
		for j := 0; j <= each; j++ {
			v = append(v, oranges{A[i]*j + B[i]*(each-j), i, j})
		}
	}
	sort.Slice(v, func(i, j int) bool {
		if v[i].val == v[j].val {
			return v[i].cnt < v[j].cnt
		}
		return v[i].val < v[j].val
	})

	ans := int(1e18)
	for i := 1; i <= n; i++ {
		seen[i] = -1
	}
	bad = n
	LB, UB = 0, 0
	j := 0
	for i := 0; i < len(v); i++ {
		add(v[i].id, v[i].cnt)
		for ok() && j < len(v) {
			del(v[j].id, v[j].cnt)
			j++
		}
		if j != 0 {
			ans = min(ans, v[i].val-v[j-1].val)
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
