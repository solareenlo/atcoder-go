package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1000001010

type Node struct {
	X [4][2][4][2]int
}

func (node *Node) Init0() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 4; k++ {
				for l := 0; l < 2; l++ {
					node.X[i][j][k][l] = -INF
				}
			}
		}
	}
}

func (node *Node) Init1(v, a int) {
	node.Init0()
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				tmp := 0
				if (i >> 1) != 0 {
					tmp = a
				}
				tmp1 := 0
				if k == 0 {
					tmp1 = v
				}
				val := tmp + tmp1
				if val > 0 {
					if j == 0 {
						node.X[i][j][((i<<1)&3)|k][j^1] = max(node.X[i][j][((i<<1)&3)|k][j^1], val)
					} else {
						node.X[i][j][((i<<1)&3)|k][j^1] = max(node.X[i][j][((i<<1)&3)|k][j^1], 0)
					}
				} else {
					node.X[i][j][((i<<1)&3)|k][j] = max(node.X[i][j][((i<<1)&3)|k][j], 0)
				}
			}
		}
	}
}

func merge(a, b Node) Node {
	var ret Node
	ret.Init0()
	for i1 := 0; i1 < 4; i1++ {
		for j1 := 0; j1 < 2; j1++ {
			for k := 0; k < 4; k++ {
				for l := 0; l < 2; l++ {
					for i2 := 0; i2 < 4; i2++ {
						for j2 := 0; j2 < 2; j2++ {
							ret.X[i1][j1][i2][j2] = max(ret.X[i1][j1][i2][j2], a.X[i1][j1][k][l]+b.X[k][l][i2][j2])
						}
					}
				}
			}
		}
	}
	return ret
}

const MAX_N = 1 << 16

var seg [MAX_N*2 - 1]Node

func update(k int, v Node) {
	k += MAX_N - 1
	seg[k] = v
	for k > 0 {
		k = (k - 1) / 2
		seg[k] = merge(seg[k*2+1], seg[k*2+2])
	}
}

var N, Q int
var A [50002]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := MAX_N - 1; i < MAX_N*2-1; i++ {
		var tmp Node
		tmp.Init1(0, 0)
		seg[i] = tmp
	}
	for i := 0; i < N+2; i++ {
		var tmp Node
		if i < 2 {
			tmp.Init1(A[i], 0)
		} else {
			tmp.Init1(A[i], A[i-2])
		}
		seg[i+MAX_N-1] = tmp
	}
	for i := MAX_N - 2; i >= 0; i-- {
		seg[i] = merge(seg[i*2+1], seg[i*2+2])
	}
	fmt.Fscan(in, &Q)
	for j := 0; j < Q; j++ {
		var k, v int
		fmt.Fscan(in, &k, &v)
		k--
		A[k] = v
		var tmp0 Node
		if k < 2 {
			tmp0.Init1(A[k], 0)
			update(k, tmp0)
		} else {
			tmp0.Init1(A[k], A[k-2])
			update(k, tmp0)
		}
		var tmp1 Node
		if k+2 < 2 {
			tmp1.Init1(A[k+2], 0)
			update(k+2, tmp1)
		} else {
			tmp1.Init1(A[k+2], A[k])
			update(k+2, tmp1)
		}
		m := 0
		for i := 0; i < 2; i++ {
			m = max(m, seg[0].X[0][0][0][i])
		}
		fmt.Fprintln(out, m)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
