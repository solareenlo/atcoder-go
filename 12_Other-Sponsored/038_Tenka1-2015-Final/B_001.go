package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
)

const MX = 100

var N, M, K int
var G [MX][]int
var id [MX]int
var d [MX]*big.Int
var ans [MX]int
var inv [MX]int
var reccnt int

func dfs(cnt int, now *big.Int) {
	if cnt == K {
		for i := 0; i < K; i++ {
			fmt.Println(id[ans[i]])
		}
		os.Exit(0)
	}
	reccnt++
	if reccnt < 2e7 && cnt+BitCount(now) >= K {
		id := FindFirst(now)
		for id < MX {
			now.SetBit(now, id, 0)
			ans[cnt] = id
			dfs(cnt+1, and(now, d[id]))
			id = FindNext(id, now)
		}
	}
}

func FindFirst(b *big.Int) int {
	for i := 0; i < MX; i++ {
		if b.Bit(i) == 1 {
			return i
		}
	}
	return MX
}

func FindNext(i int, b *big.Int) int {
	for j := i + 1; j < MX; j++ {
		if b.Bit(j) == 1 {
			return j
		}
	}
	return MX
}

func BitCount(n *big.Int) int {
	count := 0
	for _, v := range n.Bits() {
		count += popcount(uint64(v))
	}
	return count
}

func popcount(x uint64) int {
	const (
		m1  = 0x5555555555555555
		m2  = 0x3333333333333333
		m4  = 0x0f0f0f0f0f0f0f0f
		h01 = 0x0101010101010101
	)
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return int((x * h01) >> 56)
}

func main() {
	for i := range d {
		d[i] = big.NewInt(0)
	}

	fmt.Scan(&N, &M, &K)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	for i := 0; i < N; i++ {
		id[i] = i
	}
	for i := 0; i < N; i++ {
		inv[id[i]] = i
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			d[inv[i]].SetBit(d[inv[i]], j, 1)
		}
		for _, j := range G[i] {
			d[inv[i]].SetBit(d[inv[i]], inv[j], 0)
		}
	}
	fst := big.NewInt(0)
	for i := 0; i < N; i++ {
		fst.SetBit(fst, i, 1)
	}
	reccnt = 0
	dfs(0, fst)

	sort.Slice(id[:N], func(i, j int) bool {
		return len(G[id[i]])/3 < len(G[id[j]])/3
	})
	for i := 0; i < N; i++ {
		inv[id[i]] = i
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			d[inv[i]].SetBit(d[inv[i]], j, 1)
		}
		for _, j := range G[i] {
			d[inv[i]].SetBit(d[inv[i]], inv[j], 0)
		}
	}
	for i := 0; i < N; i++ {
		fst.SetBit(fst, i, 1)
	}
	reccnt = 0
	dfs(0, fst)
}

func lsh(A *big.Int, x uint) *big.Int {
	return new(big.Int).Lsh(A, x)
}

func rsh(A *big.Int, x uint) *big.Int {
	return new(big.Int).Rsh(A, x)
}

func and(A, B *big.Int) *big.Int {
	return new(big.Int).And(A, B)
}

func or(A, B *big.Int) *big.Int {
	return new(big.Int).Or(A, B)
}

func xor(A, B *big.Int) *big.Int {
	return new(big.Int).Xor(A, B)
}
