package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, Q int
	fmt.Fscan(in, &N, &M, &Q)

	A := make([]int, M)
	B := make([]int, M)
	S := make([]int, M)
	T := make([]int, M)
	bus := make([][]pair, N+1)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &A[i], &B[i], &S[i], &T[i])
		bus[A[i]] = append(bus[A[i]], pair{S[i], i})
	}

	for _, p := range bus {
		sort.Slice(p, func(i, j int) bool {
			return p[i].x < p[j].x || (p[i].x == p[j].x && p[i].y < p[j].y)
		})
	}

	table := make([][]int, 20)
	for i := range table {
		table[i] = make([]int, M)
	}
	for i := 0; i < M; i++ {
		itr := lowerBound(bus[B[i]], pair{T[i], -1})
		if itr == -1 {
			table[0][i] = i
		} else {
			table[0][i] = bus[B[i]][itr].y
		}
	}

	for i := 1; i < 20; i++ {
		for j := 0; j < M; j++ {
			table[i][j] = table[i-1][table[i-1][j]]
		}
	}

	for l := 0; l < Q; l++ {
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		itr := lowerBound(bus[y], pair{x, -1})
		if itr == -1 {
			fmt.Println(y)
			continue
		}
		now := bus[y][itr].y
		if z <= S[now] {
			fmt.Println(y)
			continue
		}
		if z <= T[now] {
			fmt.Println(A[now], B[now])
			continue
		}
		for i := 19; i >= 0; i-- {
			next := table[i][now]
			if T[next] < z {
				now = next
			}
		}
		itr = lowerBound(bus[B[now]], pair{T[now], -1})
		if itr == -1 {
			fmt.Println(B[now])
			continue
		}
		next := bus[B[now]][itr].y
		if z <= S[next] {
			fmt.Println(B[now])
			continue
		}
		fmt.Println(A[next], B[next])
	}
}

type pair struct{ x, y int }

func lowerBound(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x.x || (a[i].x == x.x && a[i].y >= x.y)
	})
	if idx < len(a) && a[idx].x >= x.x {
		return idx
	}
	return -1
}
