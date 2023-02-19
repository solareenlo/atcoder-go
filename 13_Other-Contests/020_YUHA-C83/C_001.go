package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	os, ids []int
	height  int
}

var fa, leaf []int
var l, r []int
var e []node

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100010

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)
	leaf = make([]int, N)
	fa = make([]int, N)
	for i := 1; i <= n; i++ {
		leaf[i] = i
		fa[i] = i
	}
	l = make([]int, N)
	r = make([]int, N)
	for i := 1; i <= n; i++ {
		r[i] = i
		l[i] = i
	}
	e = make([]node, N)
	for i := 0; i < m; i++ {
		var y, l int
		fmt.Fscan(in, &y, &l)
		e[i].height = y
		for j := 1; j <= l; j++ {
			var k int
			fmt.Fscan(in, &k)
			e[i].os = append(e[i].os, k)
		}
	}
	tmp := e[:m]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].height > tmp[j].height
	})
	for i := 0; i < m; i++ {
		for j := 0; j < len(e[i].os); j++ {
			e[i].ids = append(e[i].ids, getfa(e[i].os[j]))
		}
		sort.Ints(e[i].ids)
		vec := e[i].ids
		if len(vec) == 0 {
			continue
		}
		for j := len(vec) - 1; j >= 1; j-- {
			l[vec[j]] = leaf[getfa(vec[j-1])]
			r[leaf[getfa(vec[j-1])]] = vec[j]
			merge(vec[j], vec[j-1])
		}
	}
	res := make([]int, 0)
	for i := 1; i <= n; i++ {
		if l[i] != i {
			continue
		}
		x := i
		for {
			res = append(res, x)
			if x == r[x] {
				break
			}
			x = r[x]
		}
	}
	for q > 0 {
		q--
		var x int
		fmt.Fscan(in, &x)
		fmt.Println(res[x-1])
	}
}

func getfa(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = getfa(fa[x])
	return fa[x]
}

func merge(x, y int) {
	x = getfa(x)
	y = getfa(y)
	fa[x] = y
	leaf[y] = leaf[x]
}
