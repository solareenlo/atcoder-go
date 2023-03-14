package main

import (
	"bufio"
	"fmt"
	"os"
)

const ML = 100100

type P struct {
	x, y int
}

var g [][]P

func main() {
	in := bufio.NewReader(os.Stdin)

	var l, n int
	fmt.Fscan(in, &l, &n)
	g = make([][]P, ML)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		z := l - a
		y := b - 1
		x := l - 1 - z - y
		x++
		y++
		z++
		g[z] = append(g[z], P{x, y})
	}
	st := NewStarrySkyTree()
	st.init()
	sm := l * (l + 1) * (l + 2) / 6
	for i := l; i >= 0; i-- {
		sm -= st.sum(0, l, 1)
		for _, p := range g[i] {
			x := p.x
			y := p.y
			a := -1
			b := l
			for b-a > 1 {
				md := (a + b) / 2
				if st.get(md) < y {
					b = md
				} else {
					a = md
				}
			}
			if x <= b {
				continue
			}
			st.set(b, x, y, 1)
		}
	}
	fmt.Println(sm)
}

const S = 17
const N = 1 << S

type Node struct {
	d, sm, lz int
}

type StarrySkyTree struct {
	seg []Node
	sz  []int
}

func NewStarrySkyTree() *StarrySkyTree {
	tree := new(StarrySkyTree)
	tree.seg = make([]Node, 2*N)
	tree.sz = make([]int, 2*N)
	for i := 2*N - 1; i >= N; i-- {
		tree.sz[i] = 1
	}
	for i := N - 1; i >= 1; i-- {
		tree.sz[i] = tree.sz[i*2] + tree.sz[i*2+1]
	}
	return tree
}

func (tree *StarrySkyTree) init() {
	for i := 1; i < 2*N; i++ {
		tree.seg[i] = Node{0, 0, -1}
	}
}

func (tree *StarrySkyTree) lzdata(k, x int) {
	tree.seg[k].lz = x
	tree.seg[k].sm = x * tree.sz[k]
	tree.seg[k].d = x
}

func (tree *StarrySkyTree) push(k int) {
	if tree.seg[k].lz != -1 {
		tree.lzdata(k*2, tree.seg[k].lz)
		tree.lzdata(k*2+1, tree.seg[k].lz)
		tree.seg[k].lz = -1
	}
}

func (tree *StarrySkyTree) update(k int) {
	tree.seg[k].sm = tree.seg[k*2].sm + tree.seg[k*2+1].sm
}

func (tree *StarrySkyTree) get(k int) int {
	k += N
	for i := S; i > 0; i-- {
		tree.push(k >> i)
	}
	return tree.seg[k].d
}

func (tree *StarrySkyTree) set(a, b, x, k int) {
	if tree.sz[k] <= a || b <= 0 {
		return
	}
	if a <= 0 && tree.sz[k] <= b {
		tree.lzdata(k, x)
		return
	}
	tree.push(k)
	tree.set(a, b, x, k*2)
	tree.set(a-tree.sz[k]/2, b-tree.sz[k]/2, x, k*2+1)
	tree.update(k)
}

func (tree *StarrySkyTree) sum(a, b, k int) int {
	if tree.sz[k] <= a || b <= 0 {
		return 0
	}
	if a <= 0 && tree.sz[k] <= b {
		return tree.seg[k].sm
	}
	tree.push(k)
	res := 0
	res += tree.sum(a, b, k*2)
	res += tree.sum(a-tree.sz[k]/2, b-tree.sz[k]/2, k*2+1)
	return res
}
