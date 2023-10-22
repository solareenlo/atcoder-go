package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 1000005

var n, q, x, y, tot, cnt, qwq, rt, rt2, sum, g, r int
var d, s, son, l, qaq [N]int
var to [N][]int

func prework(i, fa int) {
	l1, l2 := 0, 0
	for _, j := range to[i] {
		if j == fa {
			continue
		}
		prework(j, i)
		if l[j]+1 > l1 {
			l2 = l1
			l1 = l[j] + 1
		} else if l[j]+1 > l2 {
			l2 = l[j] + 1
		}
	}
	qwq = max(qwq, l1+l2)
	l[i] = l1
}

func find(i, fa, Len int) {
	if qwq%2 == 1 && l[i]+Len == qwq && l[i] == Len-1 {
		rt = i
		rt2 = fa
		return
	}
	l1, l2 := Len, 0
	for _, j := range to[i] {
		if j == fa {
			continue
		}
		if l[j]+1 > l1 {
			l2 = l1
			l1 = l[j] + 1
		} else if l[j]+1 > l2 {
			l2 = l[j] + 1
		}
	}
	if qwq%2 == 0 && l1+l2 == qwq && l1 == l2 {
		rt = i
		return
	}
	for _, j := range to[i] {
		if j == fa {
			continue
		}
		if l[j]+1 == l1 {
			find(j, i, l2+1)
		} else {
			find(j, i, l1+1)
		}
		if rt != 0 {
			return
		}
	}
}

func dfs(i, fa, dep int) {
	s[i] = 0
	if d[i] == 1 {
		s[i] = 1
		sum += qwq/2 - dep
	}
	for _, j := range to[i] {
		if j == fa {
			continue
		}
		dfs(j, i, dep+1)
		s[i] += s[j]
		if s[son[i]] < s[j] {
			son[i] = j
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		to[x] = append(to[x], y)
		to[y] = append(to[y], x)
		d[x]++
		d[y]++
	}
	for i := 1; i <= n; i++ {
		if d[i] == 1 {
			tot++
		}
	}
	prework(1, 0)
	find(1, 0, 0)
	if rt2 != 0 {
		dfs(rt, rt2, 0)
		dfs(rt2, rt, 0)
		if s[rt] < s[rt2] {
			rt = rt2
		}
		cnt++
		qaq[cnt] = s[rt]
	} else {
		dfs(rt, 0, 0)
	}
	for s[son[rt]]*2 >= tot {
		cnt++
		qaq[cnt] = s[son[rt]]
		cnt++
		qaq[cnt] = s[son[rt]]
		rt = son[rt]
	}
	for i := 2; i <= cnt; i++ {
		qaq[i] += qaq[i-1]
	}
	g = sum
	dfs(rt, 0, 0)
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		fmt.Fscan(in, &r)
		if r <= g {
			fmt.Println(qwq)
			continue
		}
		r -= g
		p := lowerBound(qaq[1:cnt+1], r) + 1
		if p <= cnt {
			fmt.Println(qwq + p)
		} else {
			r -= qaq[cnt]
			tmp0, tmp1 := 0, 0
			if r%tot > 0 {
				tmp0 = 1
			}
			if r%tot > s[son[rt]] {
				tmp1 = 1
			}
			fmt.Println(qwq + cnt + r/tot*2 + tmp0 + tmp1)
		}
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
