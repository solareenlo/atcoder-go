package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct{ x, y int }

var (
	f   = make([][]pair, 200005)
	ch  = [200005][2]int{}
	val = [200005][2]int{}
	lim int
)

func dfs(x int) {
	f[x] = f[x][:0]
	if ch[x][0] == 0 {
		f[x] = append(f[x], pair{0, 0})
		return
	}
	for i := 0; i < 2; i++ {
		dfs(ch[x][i])
	}
	if len(f[ch[x][0]]) == 0 || len(f[ch[x][1]]) == 0 {
		return
	}
	v := make([]pair, 0)
	for k := 0; k < 2; k++ {
		a := ch[x][k]
		b := ch[x][k^1]
		j := 0
		for _, i := range f[a] {
			for j < len(f[b]) && i.y+f[b][j].x <= lim-val[x][0]-val[x][1] {
				j++
			}
			if j != 0 {
				v = append(v, pair{i.x + val[x][k], f[b][j-1].y + val[x][k^1]})
			}
		}
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].x < v[j].x
	})
	for _, i := range v {
		if len(f[x]) == 0 || f[x][len(f[x])-1].y > i.y {
			f[x] = append(f[x], i)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 2; i <= n; i++ {
		var x, v int
		fmt.Fscan(in, &x, &v)
		if ch[x][0] == 0 {
			ch[x][0] = i
			val[x][0] = v
		} else {
			ch[x][1] = i
			val[x][1] = v
		}
	}

	l, r := 0, 1<<60
	ans := 0
	for l <= r {
		mid := (l + r) / 2
		lim = mid
		dfs(1)
		if len(f[1]) != 0 {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	fmt.Println(ans)
}
