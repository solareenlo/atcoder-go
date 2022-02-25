package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n  int
	k  int
	a      = [200200]int{}
	g      = [200200]int{}
	f      = [200200]int{}
	re     = [200200]int{}
	tr     = [200200]int{}
	t  int = 0
	st     = [200200]seg{}
)

type seg struct{ w, l, r int }

func add(x, c int) {
	for x := n - x + 1; x <= n; x += x & -x {
		tr[x] += c
	}
}

func qry(x int) int {
	res := 0
	for x := n - x + 1; x >= 1; x -= x & -x {
		res += tr[x]
	}
	return res
}

func find(v int) int {
	nw := 0
	cnt := 0
	for j := 18; j >= 0; j-- {
		if nw+(1<<j) <= n && cnt+tr[nw+(1<<j)] < v {
			nw += (1 << j)
			cnt += tr[nw]
		}
	}
	return n - nw
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &k)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		re[a[i]] = i
		tmp := n - upperBound(g[1:n+1], a[i]) + 1
		if tmp == 1 && i != 1 {
			continue
		}
		g[n-tmp+1] = max(g[n-tmp+1], a[i])
		f[i] = tmp
	}

	mx := 0
	pos := 0
	for i := 1; i <= n; i++ {
		if f[re[i]] <= 0 {
			continue
		}
		if k-f[re[i]] <= 0 {
			mx = n
			pos = 0
			break
		} else {
			tmp := qry(1)
			if tmp+f[re[i]] >= k {
				tmp = find(k - f[re[i]])
				if tmp > re[i] && tmp-1 > mx {
					mx = tmp - 1
					pos = i
				}
			}
		}
		add(re[i], 1)
	}

	if mx == 0 {
		pos = k + 1
	}
	lst := n + 1
	cnt := k - f[re[pos]]
	for i := n; i >= 1 && cnt != 0; i-- {
		if a[i] < pos {
			t++
			st[t] = seg{a[i], i, lst - 1}
			lst = i
			cnt--
		}
	}

	for i := 1; i <= mx; i++ {
		fmt.Print(a[i], " ")
	}
	tmp1 := st[1 : t+1]
	sort.Slice(tmp1, func(i, j int) bool {
		return tmp1[i].w > tmp1[j].w
	})
	for j := 1; j <= t; j++ {
		for i := st[j].l; i <= st[j].r; i++ {
			fmt.Print(a[i], " ")
		}
	}
	fmt.Fprintln(out)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
