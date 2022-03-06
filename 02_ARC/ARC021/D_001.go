package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"math/bits"
	"os"
	"sort"
)

var in = bufio.NewReader(os.Stdin)

const N = 5000
const M = 200
const THRES = 83

var (
	P      = [N][M]int{}
	A      = [N]float64{}
	parent = [N]int{}
	bit    = make([]big.Int, N)
)

func INIT() {
	var x, y, z uint32
	x = 123456789
	y = 362436069
	z = 521288629
	var w uint32
	fmt.Fscan(in, &w)
	for i := 0; i < N; i++ {
		sum := 0
		for j := 0; j < M; j++ {
			var t uint32
			t = x ^ (x << 11)
			x = y
			y = z
			z = w
			w = (w ^ (w >> 19)) ^ (t ^ (t >> 8))
			var v int
			v = int(w%100000) - 50000
			if v < 0 {
				P[i][j] = v
			} else {
				P[i][j] = v + 1
			}
			sum += P[i][j] * P[i][j]
			if v > 0 {
				bit[i].SetBit(&bit[i], j, 1)
			}
		}
		A[i] = math.Sqrt(float64(sum))
	}
	for i := range parent {
		parent[i] = -1
	}
}

func cost(i, j int) float64 {
	sum := 0
	for k := 0; k < M; k++ {
		sum += P[i][k] * P[j][k]
	}
	return 1.0 - float64(sum)/(A[i]*A[j])
}

func find(u int) int {
	if parent[u] < 0 {
		return u
	}
	parent[u] = find(parent[u])
	return parent[u]
}

func merge(x, y int) bool {
	x = find(x)
	y = find(y)
	if x == y {
		return false
	}
	if parent[x] > parent[y] {
		x, y = y, x
	}
	parent[x] += parent[y]
	parent[y] = x
	return true
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	INIT()

	type pair struct {
		f float64
		k int
	}
	cand := make([]pair, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			tmp := bit[i].Xor(&bit[i], &bit[j])
			if BitCount(tmp) > THRES {
				continue
			}
			cand = append(cand, pair{cost(i, j), i*N + j})
		}
	}
	sort.Slice(cand, func(i, j int) bool {
		return cand[i].f < cand[j].f
	})

	for i := range cand {
		j := cand[i].k / N
		k := cand[i].k % N
		if merge(j, k) {
			fmt.Fprintln(out, j+1, k+1)
		}
	}
}

func BitCount(n *big.Int) int {
	count := 0
	for _, v := range n.Bits() {
		count += bits.OnesCount(uint(v))
	}
	return count
}
