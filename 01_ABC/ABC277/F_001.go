package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)
	A := make([][]int, H)
	for i := range A {
		A[i] = make([]int, W)
	}
	minA := make([]int, H)
	for i := range minA {
		minA[i] = int(1e18)
	}
	maxA := make([]int, H)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &A[i][j])
			if A[i][j] > 0 {
				minA[i] = min(minA[i], A[i][j])
			}
			maxA[i] = max(maxA[i], A[i][j])
		}
	}

	minMax := make([]pair, 0)
	for i := 0; i < H; i++ {
		if maxA[i] <= 0 {
			continue
		}
		minMax = append(minMax, pair{minA[i], maxA[i]})
	}
	sort.Slice(minMax, func(i, j int) bool {
		if minMax[i].x == minMax[j].x {
			return minMax[i].y < minMax[j].y
		}
		return minMax[i].x < minMax[j].x
	})

	for i := 0; i < len(minMax)-1; i++ {
		if minMax[i].y > minMax[i+1].x {
			fmt.Println("No")
			return
		}
	}

	num := W + H*W
	G := make([][]int, num)
	for i := range G {
		G[i] = make([]int, 0)
	}
	d := make([]int, num)
	n := W

	var addEdge func(int, int)
	addEdge = func(u, v int) {
		G[u] = append(G[u], v)
		d[v]++
	}

	for i := 0; i < H; i++ {
		list := make([]pair, W)
		for j := 0; j < W; j++ {
			list[j].x = A[i][j]
			list[j].y = j
		}
		sort.Slice(list, func(i, j int) bool {
			if list[i].x == list[j].x {
				return list[i].y < list[j].y
			}
			return list[i].x < list[j].x
		})
		current := W - 1
		for j := W - 1; j >= 0; j-- {
			if list[j].x == 0 {
				break
			}
			if list[j].x == list[current].x {
				addEdge(n, list[j].y)
				continue
			}
			n++
			addEdge(n, list[j].y)
			for k := current; k > j; k-- {
				addEdge(list[k].y, n)
			}
			current = j
		}
		n++
	}

	q := make([]int, 0)
	for i := 0; i < num; i++ {
		if d[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range G[u] {
			d[v]--
			if d[v] <= 0 {
				q = append(q, v)
			}
		}
	}

	for i := 0; i < num; i++ {
		if d[i] > 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
