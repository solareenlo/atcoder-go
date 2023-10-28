package main

import (
	"bufio"
	"fmt"
	"os"
)

type Matrix struct {
	a [2][2]int
}

const MOD = 1000000007
const BASE = 30

func (mat *Matrix) init() {
	for i := 0; i < 4; i++ {
		mat.a[i/2][i%2] = 0
	}
}

func Mul(A, B Matrix) Matrix {
	var C Matrix
	C.init()
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				C.a[i][j] += A.a[i][k] * B.a[k][j]
				C.a[i][j] %= MOD
			}
		}
	}
	return C
}

type Node struct {
	X      Matrix
	cl, cr int
}

type SegmentTree struct {
	G   []Node
	pre [33]Matrix
	I   Matrix
}

func (seg *SegmentTree) init() {
	seg.pre[0].a[0][0] = 1
	seg.pre[0].a[0][1] = 1
	seg.pre[0].a[1][0] = 1
	seg.pre[0].a[1][1] = 0
	for i := 1; i <= 32; i++ {
		seg.pre[i] = Mul(seg.pre[i-1], seg.pre[i-1])
	}
	seg.I.a[0][0] = 1
	seg.I.a[0][1] = 0
	seg.I.a[1][0] = 0
	seg.I.a[1][1] = 1
	seg.G = append(seg.G, Node{seg.pre[BASE], -1, -1})
}

func (seg *SegmentTree) query_(l, r, a, b, dep, u int) Matrix {
	if b <= l || r <= a {
		return seg.I
	}
	if l <= a && b <= r {
		if u == -1 {
			return seg.pre[dep]
		}
		return seg.G[u].X
	}
	pl, pr := -1, -1
	if u != -1 {
		pl = seg.G[u].cl
		pr = seg.G[u].cr
	}
	J1 := seg.query_(l, r, a, (a+b)>>1, dep-1, pl)
	J2 := seg.query_(l, r, (a+b)>>1, b, dep-1, pr)
	return Mul(J1, J2)
}

func (seg *SegmentTree) query(l, r int) Matrix {
	return seg.query_(l, r, 0, (1 << BASE), BASE, 0)
}

func (seg *SegmentTree) update(pos int, A Matrix) {
	cx := 0
	L := make([]int, 0)
	L = append(L, cx)
	for i := BASE - 1; i >= 0; i-- {
		if (pos/(1<<i))%2 == 0 {
			if seg.G[cx].cl == -1 {
				seg.G = append(seg.G, Node{seg.pre[i], -1, -1})
				seg.G[cx].cl = len(seg.G) - 1
			}
			cx = seg.G[cx].cl
		} else {
			if seg.G[cx].cr == -1 {
				seg.G = append(seg.G, Node{seg.pre[i], -1, -1})
				seg.G[cx].cr = len(seg.G) - 1
			}
			cx = seg.G[cx].cr
		}
		L = append(L, cx)
	}

	seg.G[cx].X = A
	for i := len(L) - 2; i >= 0; i-- {
		J1 := seg.pre[BASE-1-i]
		if seg.G[L[i]].cl >= 0 {
			J1 = seg.G[seg.G[L[i]].cl].X
		}
		J2 := seg.pre[BASE-1-i]
		if seg.G[L[i]].cr >= 0 {
			J2 = seg.G[seg.G[L[i]].cr].X
		}
		seg.G[L[i]].X = Mul(J1, J2)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)
	var X SegmentTree
	X.init()

	Map := make(map[int]int)
	for i := 1; i <= Q; i++ {
		var p int
		fmt.Fscan(in, &p)
		if p == 1 {
			var x int
			fmt.Fscan(in, &x)
			if Map[x] == 1 {
				X.update(x, Matrix{[2][2]int{{1, 1}, {1, 0}}})
			}
			if Map[x] == 0 {
				X.update(x, Matrix{[2][2]int{{0, 1}, {0, 0}}})
			}
			Map[x] ^= 1
		}
		if p == 2 {
			var l, r int
			fmt.Fscan(in, &l, &r)
			if Map[l] == 1 || Map[r] == 1 {
				fmt.Fprintln(out, 0)
			} else {
				G := X.query(l+1, r+1)
				fmt.Fprintln(out, G.a[0][0])
			}
		}
	}
}
