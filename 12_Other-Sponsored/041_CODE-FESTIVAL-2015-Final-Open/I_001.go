package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 100005

type bit struct {
	tree []int
}

func newBit(size int) *bit {
	return &bit{tree: make([]int, size)}
}

func (b *bit) add(x, v int) {
	x++
	for x <= len(b.tree) {
		b.tree[x] += v
		x += x & -x
	}
}

func (b *bit) sum(x int) int {
	x++
	ret := 0
	for x > 0 {
		ret += b.tree[x]
		x -= x & -x
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	dep := make([]int, n+1)
	par := make([]int, n+1)
	gph := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &dep[i])
	}

	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &par[i])
		gph[par[i]] = append(gph[par[i]], i)
		dep[i] += dep[par[i]]
	}

	lis := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		lis = append(lis, dep[i])
	}
	sort.Ints(lis)
	uniqueDep := make([]int, 0, len(lis))
	uniqueDep = append(uniqueDep, lis[0])
	for i := 1; i < len(lis); i++ {
		if lis[i] != lis[i-1] {
			uniqueDep = append(uniqueDep, lis[i])
		}
	}

	dap := make([]int, n+1)
	mxson := make([]int, n+1)
	bit := newBit(N)

	for i := 1; i <= n; i++ {
		dep[i] = sort.SearchInts(uniqueDep, dep[i])
	}

	for i := n; i > 0; i-- {
		mxson[i] = dep[i]
		for _, j := range gph[i] {
			mxson[i] = max(mxson[i], mxson[j])
		}
	}

	dap = make([]int, len(uniqueDep))
	for i := range dap {
		dap[i] = int(1e9)
	}

	dfs(1, dap, bit, mxson, gph, dep)

	var q int
	fmt.Fscan(in, &q)
	for t := 0; t < q; t++ {
		var x int
		fmt.Fscan(in, &x)
		pos := sort.SearchInts(uniqueDep, x)
		if pos == len(uniqueDep) || uniqueDep[pos] != x || dap[pos] > int(1e9) {
			fmt.Fprintln(out, "-1")
		} else {
			fmt.Fprintln(out, dap[pos])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(x int, dap []int, bit *bit, mxson []int, gph [][]int, dep []int) {
	for _, i := range gph[x] {
		bit.add(mxson[i], 1)
	}
	dap[dep[x]] = min(dap[dep[x]], bit.sum(len(dep)-1)-bit.sum(dep[x]))
	for _, i := range gph[x] {
		bit.add(mxson[i], -1)
		dfs(i, dap, bit, mxson, gph, dep)
		bit.add(mxson[i], 1)
	}
	for _, i := range gph[x] {
		bit.add(mxson[i], -1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
