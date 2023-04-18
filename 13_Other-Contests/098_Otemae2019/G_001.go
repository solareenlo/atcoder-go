package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

const ma2 = (1 << 11) - 1

var x, y, ima int
var zaatuX, zaatuY [2010]int
var hikari, DD [2010][2010]int
var PQ *Heap

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	var T [1001]bool
	var A, B, C, D [1001]int
	sortX := make([]int, 2010)
	sortY := make([]int, 2010)
	var P, Q, R, S int
	fmt.Fscan(in, &P, &Q, &R, &S)
	sortX[0] = (P << 32)
	sortY[0] = (Q << 32)
	sortX[1] = (R << 32) ^ 1
	sortY[1] = (S << 32) ^ 1
	k := 2
	for i := 1; i < M+1; i++ {
		var tmp int
		fmt.Fscan(in, &tmp)
		if tmp-1 != 0 {
			T[i] = true
		} else {
			T[i] = false
		}
		var A, B, C, D int
		fmt.Fscan(in, &A, &B, &C, &D)
		sortX[k] = (A << 32) ^ k
		sortX[k+1] = ((B << 32) ^ k) + 1
		sortY[k] = (C << 32) ^ k
		sortY[k+1] = ((D << 32) ^ k) + 1
		k += 2
	}

	tmpX := sortX[:M*2+2]
	sort.Ints(tmpX)
	tmpY := sortY[:M*2+2]
	sort.Ints(tmpY)

	maeX := 0
	maeY := 0
	const ma = (1 << 32) - 1
	for i := 0; i < M*2+2; i++ {
		tmpX := sortX[i] >> 32
		if tmpX != maeX {
			x++
			zaatuX[x] = tmpX
			maeX = tmpX
		}
		tmp := sortX[i] & ma
		if tmp > 1 {
			if (tmp & 1) != 0 {
				B[tmp>>1] = x
			} else {
				A[tmp>>1] = x
			}
		} else {
			if (tmp & 1) != 0 {
				R = x
			} else {
				P = x
			}
		}
		tmpY := sortY[i] >> 32
		if tmpY != maeY {
			y++
			zaatuY[y] = tmpY
			maeY = tmpY
		}
		tmp = sortY[i] & ma
		if tmp > 1 {
			if (tmp & 1) != 0 {
				D[tmp>>1] = y
			} else {
				C[tmp>>1] = y
			}
		} else {
			if (tmp & 1) != 0 {
				S = y
			} else {
				Q = y
			}
		}
	}

	const yy = 1 << 11
	for i := 1; i < M+1; i++ {
		if T[i] {
			hikari[A[i]][C[i]] += yy
			hikari[A[i]][D[i]+1] -= yy
			hikari[B[i]][C[i]] -= yy
			hikari[B[i]][D[i]+1] += yy
		} else {
			hikari[A[i]][C[i]] += 1
			hikari[A[i]][D[i]] -= 1
			hikari[B[i]+1][C[i]] -= 1
			hikari[B[i]+1][D[i]] += 1
		}
	}
	for i := 1; i < x+1; i++ {
		for j := 1; j < y+1; j++ {
			hikari[i][j] += hikari[i-1][j] + hikari[i][j-1] - hikari[i-1][j-1]
		}
	}
	for i := 1; i < x+1; i++ {
		for j := 1; j < y+1; j++ {
			DD[i][j] = 2e9
		}
	}

	DD[P][Q] = 0
	PQ = &Heap{}
	heap.Push(PQ, (P<<11)+Q)

	for PQ.Len() > 0 {
		p := heap.Pop(PQ).(int)
		i := (p >> 11) & ma2
		j := p & ma2
		d := p >> 32
		if DD[i][j] < d {
			continue
		}
		ima = d
		Warp(i, j)
	}

	fmt.Println(DD[R][S])
}

func Warp(i, j int) {
	if i > 1 {
		tmp0 := zaatuX[i] - zaatuX[i-1]
		if (hikari[i-1][j] >> 11) != 0 {
			tmp0 = 0
		}
		tmp := ima + tmp0
		if DD[i-1][j] > tmp {
			DD[i-1][j] = tmp
			heap.Push(PQ, (DD[i-1][j]<<32)+((i-1)<<11)+j)
		}
	}
	if i < x {
		tmp0 := zaatuX[i+1] - zaatuX[i]
		if (hikari[i][j] >> 11) != 0 {
			tmp0 = 0
		}
		tmp := ima + tmp0
		if DD[i+1][j] > tmp {
			DD[i+1][j] = tmp
			heap.Push(PQ, (DD[i+1][j]<<32)+((i+1)<<11)+j)
		}
	}
	if j > 1 {
		tmp0 := zaatuY[j] - zaatuY[j-1]
		if (hikari[i][j-1] & ma2) != 0 {
			tmp0 = 0
		}
		tmp := ima + tmp0
		if DD[i][j-1] > tmp {
			DD[i][j-1] = tmp
			heap.Push(PQ, (DD[i][j-1]<<32)+(i<<11)+j-1)
		}
	}
	if j < y {
		tmp0 := zaatuY[j+1] - zaatuY[j]
		if (hikari[i][j] & ma2) != 0 {
			tmp0 = 0
		}
		tmp := ima + tmp0
		if DD[i][j+1] > tmp {
			DD[i][j+1] = tmp
			heap.Push(PQ, (DD[i][j+1]<<32)+(i<<11)+j+1)
		}
	}
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
