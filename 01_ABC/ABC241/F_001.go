package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 110000
const N4 = 410000

type node struct{ x, y, id int }

var (
	a        = make([]node, N)
	b        = make([]node, N)
	st       node
	ed       node
	cmp_type int
)

func pd(x, y node, Type int) bool {
	if Type == 0 {
		if x.x == y.x {
			return x.y < y.y
		} else {
			return x.x < y.x
		}
	} else {
		if x.y == y.y {
			return x.x < y.x
		} else {
			return x.y < y.y
		}
	}
}

func cmp(x, y node) bool { return pd(x, y, cmp_type) }

func pdd(x int, y node, Type int) bool {
	if Type == 0 {
		return pd(a[x], y, Type)
	} else {
		return pd(b[x], y, Type)
	}
}

var (
	n    int
	h    = make([]int, N4)
	list = make([]node, N4)
	val  = make([]int, N4)
	head int
	tail int
)

func findpos(Type int, x node) int {
	l := 1
	r := n
	ans := 0
	for l <= r {
		mid := (l + r) >> 1
		if pdd(mid, x, Type) == true {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

var (
	dx = [5]int{0, 0, 1, 0, -1}
	dy = [5]int{0, 1, 0, -1, 0}
	H  int
	W  int
)

func id(x, y int) int { return (x-1)*4 + y }

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &H, &W, &n)
	fmt.Fscan(in, &st.x, &st.y, &ed.x, &ed.y)

	a = make([]node, n+1)
	b = make([]node, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].x, &a[i].y)
		a[i].id = i
		b[i] = a[i]
	}
	cmp_type = 0
	sort.Slice(a, func(i, j int) bool {
		if a[i].x == a[j].x {
			return a[i].y < a[j].y
		} else {
			return a[i].x < a[j].x
		}
	})
	cmp_type = 1
	sort.Slice(b, func(i, j int) bool {
		if b[i].y == b[j].y {
			return b[i].x < b[j].x
		} else {
			return b[i].y < b[j].y
		}
	})

	head = 1
	tail = 1
	list[head] = st
	val[1] = 1
	for i := 1; i <= n; i++ {
		for t := 1; t <= 4; t++ {
			x := node{a[i].x + dx[t], a[i].y + dy[t], 0}
			if x.x == st.x && x.y == st.y {
				h[id(a[i].id, t)] = 1
				break
			}
		}
	}

	for head <= tail {
		zhi := val[head]
		x := list[head]
		head++
		now := findpos(0, x)
		if now != 0 && a[now].x == x.x && a[now].y <= x.y && h[id(a[now].id, 1)] == 0 {
			tail++
			list[tail] = node{a[now].x + dx[1], a[now].y + dy[1], 0}
			h[id(a[now].id, 1)] = zhi + 1
			val[tail] = zhi + 1
		}
		now++
		if now <= n && a[now].x == x.x && a[now].y >= x.y && h[id(a[now].id, 3)] == 0 {
			tail++
			list[tail] = node{a[now].x + dx[3], a[now].y + dy[3], 0}
			h[id(a[now].id, 3)] = zhi + 1
			val[tail] = zhi + 1
		}
		now = findpos(1, x)
		if now != 0 && b[now].y == x.y && b[now].x <= x.x && h[id(b[now].id, 2)] == 0 {
			tail++
			list[tail] = node{b[now].x + dx[2], b[now].y + dy[2], 0}
			h[id(b[now].id, 2)] = zhi + 1
			val[tail] = zhi + 1
		}
		now++
		if now <= n && b[now].y == x.y && b[now].x >= x.x && h[id(b[now].id, 4)] == 0 {
			tail++
			list[tail] = node{b[now].x + dx[4], b[now].y + dy[4], 0}
			h[id(b[now].id, 4)] = zhi + 1
			val[tail] = zhi + 1
		}
	}

	ans := 999999999
	for i := 1; i <= n; i++ {
		for t := 1; t <= 4; t++ {
			x := node{a[i].x + dx[t], a[i].y + dy[t], 0}
			if x.x == ed.x && x.y == ed.y && h[id(a[i].id, t)] != 0 && h[id(a[i].id, t)] < ans {
				ans = h[id(a[i].id, t)]
			}
		}
	}
	if ans == 999999999 {
		ans = 0
	}
	fmt.Println(ans - 1)
}
