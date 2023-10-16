package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1001001001

type Node struct {
	emp bool
	a   [2][2]int
}

func newNode(val int) *Node {
	node := &Node{emp: false}
	node.a[0][0] = 0
	node.a[0][1] = INF
	node.a[1][0] = INF
	node.a[1][1] = val
	return node
}

func (n *Node) pull(l, r *Node) {
	if l.emp {
		n.emp = r.emp
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				n.a[i][j] = r.a[i][j]
			}
		}
	} else if r.emp {
		n.emp = l.emp
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				n.a[i][j] = l.a[i][j]
			}
		}
	} else {
		n.emp = false
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				n.a[i][j] = INF
			}
		}
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				for k := 0; k < 2; k++ {
					for m := 0; m < 2; m++ {
						if !(j == 0 && k == 0) {
							n.a[i][m] = min(n.a[i][m], l.a[i][j]+r.a[k][m])
						}
					}
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type SegmentTreePoint struct {
	logN int
	n    int
	ts   []*Node
}

func newSegmentTreePoint(ss []int) *SegmentTreePoint {
	n_ := len(ss)
	logN, n := 0, 1
	for n < n_ {
		logN++
		n <<= 1
	}
	ts := make([]*Node, n<<1)
	for i := range ts {
		ts[i] = new(Node)
		ts[i].emp = true
	}
	for i := 0; i < n_; i++ {
		ts[n+i] = newNode(ss[i])
	}
	build(ts, n)
	return &SegmentTreePoint{logN, n, ts}
}

func build(ts []*Node, n int) {
	for u := n - 1; u > 0; u-- {
		pull(ts, u)
	}
}

func (st *SegmentTreePoint) change(a int, s int) {
	a += st.n
	st.ts[a] = newNode(s)
	for a >>= 1; a > 0; a >>= 1 {
		pull(st.ts, a)
	}
}

func pull(ts []*Node, u int) {
	ts[u].pull(ts[u<<1], ts[u<<1|1])
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	S := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &S[i])
		S[i] *= 2
	}

	var sumS int
	for i := 0; i < N; i++ {
		sumS += S[i]
	}
	seg := newSegmentTreePoint(S)

	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var P int
		var X int
		fmt.Fscan(in, &P, &X)
		P--
		X *= 2
		sumS -= S[P]
		S[P] = X
		sumS += S[P]
		seg.change(P, X)
		f := seg.ts[1]
		ans := sumS / 2
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				if !(i == 0 && j == 0) {
					ans = min(ans, f.a[i][j])
				}
			}
		}
		fmt.Fprintln(out, ans)
	}
}
