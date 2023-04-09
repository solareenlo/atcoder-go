package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type pair struct {
		x, y int
	}

	var n, m int
	var tmp string
	fmt.Fscan(in, &n, &tmp, &m)
	s := strings.Split(tmp, "")
	d := make([]string, m)
	f := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &d[i], &f[i])
	}
	var q int
	fmt.Fscan(in, &q)
	t := make([]int, q)
	p := make([]int, q)
	var query [200000][]pair
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &t[i], &p[i])
		t[i]--
		query[t[i]] = append(query[t[i]], pair{p[i], i})
	}

	type P struct {
		x string
		y int
	}
	que := list.New()
	for i := 0; i < n; i++ {
		if que.Len() == 0 || que.Back().Value.(P).x != s[i] {
			que.PushBack(P{s[i], 1})
		} else {
			tmp := que.Back().Value.(P)
			que.Back().Value = P{tmp.x, tmp.y + 1}
		}
	}
	cnt := 0
	ans := make([]string, 200000)
	for i := 0; i < m; i++ {
		if d[i][0] == 'L' {
			cnt++
			if que.Front().Value.(P).x != f[i] {
				p := que.Front().Value.(P)
				que.Remove(que.Front())
				if que.Len() == 0 {
					que.PushFront(p)
					que.PushFront(P{f[i], 1})
				} else {
					tmp := que.Front().Value.(P)
					que.Front().Value = P{tmp.x, tmp.y + p.y + 1}
				}
			} else {
				tmp := que.Front().Value.(P)
				que.Front().Value = P{tmp.x, tmp.y + 1}
			}
		} else {
			if que.Back().Value.(P).x != f[i] {
				p := que.Back().Value.(P)
				que.Remove(que.Back())
				if que.Len() == 0 {
					que.PushBack(p)
					que.PushBack(P{f[i], 1})
				} else {
					tmp := que.Back().Value.(P)
					que.Back().Value = P{tmp.x, tmp.y + p.y + 1}
				}
			} else {
				tmp := que.Back().Value.(P)
				que.Back().Value = P{tmp.x, tmp.y + 1}
			}
		}
		for _, p := range query[i] {
			if p.x <= que.Front().Value.(P).y {
				ans[p.y] = que.Front().Value.(P).x
			} else if p.x > n+i+1-que.Back().Value.(P).y {
				ans[p.y] = que.Back().Value.(P).x
			} else {
				ans[p.y] = s[p.x-cnt-1]
			}
		}
	}
	for i := 0; i < q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
