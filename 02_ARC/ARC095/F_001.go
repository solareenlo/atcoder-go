package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100004

var (
	fa  = make([]int, N)
	dep = make([]int, N)
	E   = make([][]int, N)
)

func dfs(u, f int) {
	fa[u] = f
	dep[u] = dep[f] + 1
	for _, v := range E[u] {
		if v != f {
			dfs(v, u)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		E[x] = append(E[x], y)
		E[y] = append(E[y], x)
	}

	dfs(1, 0)
	P := 0
	for i := 1; i <= n; i++ {
		if dep[i] > dep[P] {
			P = i
		}
	}

	dfs(P, 0)
	Q := 0
	for i := 1; i <= n; i++ {
		if dep[i] > dep[Q] {
			Q = i
		}
	}

	leng := 0
	mrk := make([]int, N)
	for i := Q; i > 0; i = fa[i] {
		leng++
		mrk[i] = leng
	}

	t := make([]int, N)
	for i := 1; i <= n; i++ {
		if mrk[i] == 0 {
			if mrk[fa[i]] == 0 {
				fmt.Fprintln(out, -1)
				return
			}
			t[mrk[fa[i]]]++
		}
	}

	for i := 1; i <= leng; i++ {
		if t[i] != t[leng-i+1] {
			if t[i] > t[leng-i+1] {
				tmp := t[1 : leng+1]
				tmp = reverseOrderInt(tmp)
				for i := 0; i < leng; i++ {
					t[i+1] = tmp[i]
				}
			}
			break
		}
	}

	for i, j := 1, 0; i <= leng; i++ {
		for k := j + 2; k <= j+t[i]+1; k++ {
			fmt.Fprint(out, k, " ")
		}
		fmt.Fprint(out, j+1, " ")
		j += t[i] + 1
	}
	fmt.Fprintln(out)
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
