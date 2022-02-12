package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 1_000_000_007
const N = 210

type node struct {
	l, r  int
	delta []int
}

type Node []node
type pair struct{ x, y int }

func (a Node) Len() int { return len(a) }
func (a Node) Less(i, j int) bool {
	for k := 0; k < len(a[i].delta); k++ {
		if a[i].delta[k] != a[j].delta[k] {
			return a[i].delta[k] < a[j].delta[k]
		}
	}
	return false

}
func (a Node) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

var (
	n   int
	p   = make([]int, N)
	vec = make([]node, 0)
)

func query(l, r int) int {
	res := 0
	for i := 61; i >= 0; i-- {
		x := l >> i & 1
		y := r >> i & 1
		l -= x << i
		r -= y << i
		if x == y {
			res += x
			continue
		}
		if l == 0 {
			break
		}
		res++
		break
	}
	return res
}

func calc(x int) {
	if x < 0 {
		return
	}
	mn := p[1]
	for i := 1; i <= n; i++ {
		mn = min(mn, p[i])
	}
	if mn < 0 {
		return
	}
	A := new(node)
	A.delta = make([]int, n-1)
	A.l = max(0, mn-x)
	A.r = mn
	for i := 1; i < n; i++ {
		A.delta[i-1] = p[i] - p[i+1]
	}
	vec = append(vec, *A)
}

func update(x *int, y int) {
	*x += y
	if *x >= mod {
		*x -= mod
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var K int
	fmt.Fscan(in, &n, &K)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	b := make([]pair, N)
	for i := 0; i <= 60; i++ {
		for j := 1; j <= n; j++ {
			b[j] = pair{a[j] % (1 << i), j}
		}
		tmp := b[1 : n+1]
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i].x < tmp[j].x
		})
		for j := 1; j <= n; j++ {
			p[j] = a[j] / (1 << i)
		}
		calc(K - i - query(0, b[1].x))
		b[n+1].x = (1 << i) - 1
		for j := 1; j <= n; j++ {
			p[b[j].y]--
			if b[j].x < b[j+1].x {
				calc(K - i - query(b[j].x+1, b[j+1].x))
			}
		}
	}
	sort.Sort(Node(vec))

	cmp := func(A, B []int) bool {
		for i := 0; i < len(A); i++ {
			if A[i] != B[i] {
				return false
			}
		}
		return true
	}
	ans := 0
	j := 0
	for i := 0; i < len(vec); i = j + 1 {
		j = i
		V := make([]pair, 0)
		V = append(V, pair{vec[i].l, vec[i].r})
		for j+1 < len(vec) && cmp(vec[i].delta, vec[j+1].delta) {
			j++
			V = append(V, pair{vec[j].l, vec[j].r})
		}
		sort.Slice(V, func(i, j int) bool {
			return V[i].x < V[j].x
		})
		L := -1
		R := -1
		for _, A := range V {
			if L == -1 {
				L = A.x
				R = A.y
			} else if A.x > R {
				if L != -1 {
					update(&ans, (R-L+1)%mod)
				}
				L = A.x
				R = A.y
			} else {
				R = max(R, A.y)
			}
		}
		if L != -1 {
			update(&ans, (R-L+1)%mod)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
