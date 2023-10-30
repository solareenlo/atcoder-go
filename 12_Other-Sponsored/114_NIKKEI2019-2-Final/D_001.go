package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x int
	y uint
}

func xorshift64(x uint) uint {
	x ^= x << 13
	x ^= x >> 7
	x ^= x << 17
	return x
}

func solve_hash_a(in [][]pair, sum []uint, ter, per int) uint {
	ret := uint(231093)
	for it := range in[ter] {
		if in[ter][it].x == per {
			continue
		}
		in[ter][it].y = solve_hash_a(in, sum, in[ter][it].x, ter)
		ret += in[ter][it].y
	}
	sum[ter] = ret
	return xorshift64(ret)
}

func solve_hash_b(in [][]pair, sum []uint, ter, per int, pH uint) {
	sum[ter] += pH
	for it := range in[ter] {
		if in[ter][it].x == per {
			in[ter][it].y = pH
			continue
		}
		solve_hash_b(in, sum, in[ter][it].x, ter, xorshift64(sum[ter]-in[ter][it].y))
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	koV := make([][]pair, n+1)
	koU := make([][]pair, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(in, &v)
		v--
		koV[i] = append(koV[i], pair{v, 0})
		koV[v] = append(koV[v], pair{i, 0})
	}
	for i := 0; i < n-1; i++ {
		var u int
		fmt.Fscan(in, &u)
		u--
		koU[i] = append(koU[i], pair{u, 0})
		koU[u] = append(koU[u], pair{i, 0})
	}
	sV := make([]uint, n+1)
	sU := make([]uint, n)
	solve_hash_a(koV, sV, 0, -1)
	solve_hash_b(koV, sV, 0, -1, 0)
	solve_hash_a(koU, sU, 0, -1)
	solve_hash_b(koU, sU, 0, -1, 0)
	for i := 0; i < n; i++ {
		sU[i] = xorshift64(sU[i])
	}
	sort.Slice(sU, func(i, j int) bool {
		return sU[i] < sU[j]
	})
	for i := 0; i < n+1; i++ {
		if len(koV[i]) != 1 {
			continue
		}
		x := koV[i][0].y
		idx := lowerBound(sU, x)
		if idx < len(sU) && sU[idx] == x {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
}

func lowerBound(a []uint, x uint) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
