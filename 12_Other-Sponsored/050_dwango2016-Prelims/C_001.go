package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type tuple struct {
		a, b, c, d int
	}

	var n, m, s, t int
	fmt.Fscan(in, &n, &m, &s, &t)
	g := make([][]tuple, n)
	for i := 0; i < m; i++ {
		var l int
		fmt.Fscan(in, &l)
		path := make([]int, l)
		cost := make([]int, l-1)
		for j := 0; j < l; j++ {
			fmt.Fscan(in, &path[j])
		}
		for j := 0; j < l-1; j++ {
			fmt.Fscan(in, &cost[j])
		}
		for j := 0; j < 2; j++ {
			sum := accumulate(cost)
			for k := 0; k < l-1; k++ {
				u := path[k]
				v := path[k+1]
				w := cost[k]
				g[v] = append(g[v], tuple{u, w, path[len(path)-1], sum})
				sum -= w
			}
			reverseOrderInt(path)
			reverseOrderInt(cost)
		}
	}
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1 << 29
		}
	}
	for i := 0; i < 2; i++ {
		pq := make(PriorityQueue, 0)
		heap.Init(&pq)
		heap.Push(&pq, Node{0, t})
		dp[i][t] = 0
		for pq.Len() > 0 {
			tmp := heap.Pop(&pq).(Node)
			c := tmp.cost
			v := tmp.vertex
			for _, it := range g[v] {
				u, w, end, wsum := it.a, it.b, it.c, it.d
				var nc int
				if i == 0 {
					nc = w + c
				}
				if i == 1 {
					nc = max(w+c, wsum+dp[0][end])
				}
				if dp[i][u] > nc {
					dp[i][u] = nc
					heap.Push(&pq, Node{dp[i][u], u})
				}
			}
		}
	}
	fmt.Println(dp[1][s])
}

type Node struct {
	cost   int
	vertex int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func accumulate(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
