package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const SIZE = 4005

	var A, B, C, D, E, F [SIZE]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	vx := make([]int, 0)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		x := b - 1
		y := n - (a + c - 1)
		A[i] = x
		B[i] = y
		C[i] = c
		vx = append(vx, A[i])
		vx = append(vx, A[i]+C[i])
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		x := b - 1
		y := n - (a + c - 1)
		D[i] = x
		E[i] = y
		F[i] = c
		vx = append(vx, D[i])
		vx = append(vx, D[i]+F[i])
	}
	sort.Ints(vx)
	vx = unique(vx)
	query := make([][]P, SIZE)
	for i := 0; i < m; i++ {
		pos := lowerBound(vx, A[i])
		for vx[pos] < A[i]+C[i] {
			h := A[i] + C[i] - vx[pos+1]
			query[pos] = append(query[pos], P{B[i], h})
			pos++
		}
	}
	vec := make([][]PP, SIZE)
	vec2 := make([][]PP, SIZE)
	sum := make([][]int, SIZE)
	for i := 0; i+1 < len(vx); i++ {
		sortPAIR(query[i])
		now := 0
		w := vx[i+1] - vx[i]
		for j := 0; j < len(query[i]); {
			h := query[i][j].y
			bot := query[i][j].x
			for ; j < len(query[i]) && bot+h+1 >= query[i][j].x; j++ {
				h = max(h, query[i][j].x+query[i][j].y-bot)
			}
			Len := w
			if j < len(query[i]) {
				Len = min(Len, query[i][j].x-(bot+h))
			}
			now += w*(h+Len) - Len*(Len-1)/2
			vec[i] = append(vec[i], PP{bot + h + Len, P{h, Len}})
			vec2[i] = append(vec2[i], PP{bot + h, P{h, Len}})
			sum[i] = append(sum[i], now)
		}
	}
	for i := 0; i < q; i++ {
		ans := 0
		pos := lowerBound(vx, D[i])
		for vx[pos] < D[i]+F[i] {
			h := D[i] + F[i] - vx[pos+1]
			w := vx[pos+1] - vx[pos]
			down := lowerBoundPair(vec[pos], PP{E[i] + 1, P{-1, -1}})
			if down > 0 {
				ans -= sum[pos][down-1]
			}
			if down < len(vec[pos]) {
				p := vec[pos][down]
				h2 := p.y.x
				l2 := p.y.y
				d2 := p.x - (h2 + l2)
				if d2 < E[i] {
					if d2+h2 >= E[i] {
						ans -= w * (E[i] - d2)
					} else {
						ans -= w * h2
						Len := E[i] - (d2 + h2)
						ans -= w*Len - Len*(Len-1)/2
					}
				}
			}
			up := lowerBoundPair(vec2[pos], PP{E[i] + h + 1, P{-1, -1}})
			if up > 0 {
				ans += sum[pos][up-1]
			}
			if up < len(vec2[pos]) {
				p := vec2[pos][up]
				h2 := p.y.x
				d2 := p.x - h2
				if d2 <= E[i]+h+w-1 {
					ans += w * (w + 1) / 2
					if d2 >= E[i]+h {
						b := d2 - (E[i] + h)
						ans -= w*b - b*(b-1)/2
					} else {
						ans += w * (E[i] + h - d2)
					}
				}
			}
			pos++
		}
		all := F[i] * (F[i] + 1) / 2
		fmt.Println(all - ans)
	}
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
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

type PP struct {
	x int
	y P
}

func lowerBoundPair(a []PP, x PP) int {
	idx := sort.Search(len(a), func(i int) bool {
		if a[i].x == x.x {
			if a[i].y.x == x.y.x {
				return a[i].y.y >= x.y.y
			}
			return a[i].y.x >= x.y.x
		}
		return a[i].x >= x.x
	})
	return idx
}

type P struct {
	x, y int
}

func sortPAIR(tmp []P) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
