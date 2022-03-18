package main

import (
	"bufio"
	"fmt"
	"os"
)

type Data struct{ l, r int }

var (
	son = [140010][2]int{}
	val = [140010][2]int{}
	bd  = [140010][2]int{}
	tot int
	a   = make([]Data, 140010)
	s   = [2]int{}
	mx  int
	q   = [140010][2]Data{}
)

func DFS(x int) {
	if son[x][0] == 0 {
		tot++
		a[tot] = Data{0, 0}
		bd[x][0] = tot
		bd[x][1] = tot
		return
	}
	DFS(son[x][0])
	y := son[x][0]
	z := val[x][0]
	for i := bd[y][0]; i <= bd[y][1]; i++ {
		a[i].l += z
		a[i].r += z
	}
	if son[x][1] == 0 {
		bd[x][0] = bd[y][0]
		bd[x][1] = bd[y][1]
		return
	}
	DFS(son[x][1])
	y = son[x][1]
	z = val[x][1]
	for i := bd[y][0]; i <= bd[y][1]; i++ {
		a[i].l += z
		a[i].r += z
	}
	s[0] = 0
	s[1] = 0
	for i := 0; i < 2; i++ {
		y = son[x][i]
		z = son[x][i^1]
		for j, k := bd[y][0], bd[z][0]; j <= bd[y][1]; j++ {
			for k <= bd[z][1] && a[j].r+a[k].l <= mx {
				k++
			}
			if k == bd[z][0] {
				continue
			}
			s[i]++
			q[s[i]][i] = Data{a[j].l, a[k-1].r}
		}
	}
	y = son[x][0]
	z = son[x][1]
	i := 1
	j := 1
	k := bd[y][0] - 1
	for ; i <= s[0]; i++ {
		for j <= s[1] && q[j][1].l <= q[i][0].l {
			if k < bd[y][0] || q[j][1].r < a[k].r {
				k++
				a[k] = q[j][1]
			}
			j++
		}
		if k >= bd[y][0] && a[k].l == q[i][0].l {
			a[k].r = min(a[k].r, q[i][0].r)
		} else if k < bd[y][0] || q[i][0].r < a[k].r {
			k++
			a[k] = q[i][0]
		}
	}
	for ; j <= s[1]; j++ {
		if k < bd[y][0] || q[j][1].r < a[k].r {
			k++
			a[k] = q[j][1]
		}
	}
	bd[x][0] = bd[y][0]
	bd[x][1] = k
}

func check(mid int) bool {
	tot = 0
	mx = mid
	DFS(1)
	return bd[1][0] <= bd[1][1]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 2; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if son[x][0] == 0 {
			son[x][0] = i
			val[x][0] = y
		} else {
			son[x][1] = i
			val[x][1] = y
		}
	}

	l := -1
	r := 4200000
	for l < r-1 {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Println(r)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
