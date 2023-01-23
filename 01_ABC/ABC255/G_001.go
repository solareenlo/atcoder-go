package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

const MAXN = 200005

var M int
var sg [MAXN]int
var mx [MAXN]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, MAXN)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	b := make([]pair, MAXN)
	c := make([]int, MAXN)
	C := make([]int, MAXN<<1)
	N := 0
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		b[i].x = x
		b[i].y = x - y
		c[i] = x
		N++
		C[N] = x
		N++
		C[N] = x - y
	}

	tmp := b[1 : m+1]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
	for i := 0; i < m; i++ {
		b[1+i] = tmp[i]
	}

	sort.Ints(c[1 : m+1])
	M = len(unique(c[1 : m+1]))

	sort.Ints(C[1 : N+1])
	N = len(unique(C[1 : N+1]))

	d := make([]int, MAXN)
	p := 1
	for i := 1; i <= m; i++ {
		if b[i].x == c[p] {
			d[i] = p
		} else {
			p++
			d[i] = p
		}
	}

	fa := make([]int, MAXN<<1)
	s := make([]int, MAXN<<1)
	for i := 1; i <= N; i++ {
		fa[i] = i
		s[i] = 1
	}

	st := make([]int, MAXN)
	S := make([]int, MAXN<<1)
	for i := 1; i <= m; i++ {
		if b[i].x != b[i-1].x {
			sg[d[i]] = mx[d[i]-1] + b[i].x - b[i-1].x
			j := i
			for b[j+1].x == b[i].x {
				j++
			}
			top := 0
			for k := i; k <= j; k++ {
				x := lowerBound(C[1:N+1], b[k].y) + 1
				// fmt.Println(x)
				top++
				st[top] = fa[x]
				S[st[top]]++
			}
			p = 0
			for top != 0 {
				if S[st[top]] == s[st[top]] {
					if sg[d[i]] > getsg(C[st[top]], c) {
						sg[d[i]] = getsg(C[st[top]], c)
						p = st[top]
					}
				}
				S[st[top]] = 0
				top--
			}
			if p != 0 {
				fa[lowerBound(C[1:N+1], b[i].x)+1] = p
				s[p]++
			}
			mx[d[i]] = MAX(mx[d[i]-1]+b[i].x-b[i-1].x-1, sg[d[i]])
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans ^= getsg(a[i], c)
	}
	if ans != 0 {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getsg(x int, c []int) int {
	i := lowerBound(c[1:M+1], x) + 1
	if c[i] == x {
		return sg[i]
	}
	if i == 0 {
		return x
	}
	return mx[i-1] + x - c[i-1]
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}
