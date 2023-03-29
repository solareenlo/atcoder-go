package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &C[i])
	}
	C[0] = min(C[0], C[N-1])
	train := make([][]Train, N-1)
	for i := 0; i < M; i++ {
		var S, T, A, B int
		fmt.Fscan(in, &S, &T, &A, &B)
		train[A-1] = append(train[A-1], Train{S, T, B})
	}
	for t := range train {
		sort.Slice(train[t], func(i, j int) bool {
			return train[t][i].S < train[t][j].S
		})
	}
	valid := true
	startIdx := make([]int, N)
	startIdx[0] = 2
	for i := 0; i < N-1; i++ {
		if len(train[i]) != 0 {
			startIdx[i+1] = startIdx[i] + len(train[i]) - 1
		} else {
			valid = false
		}
	}
	if !valid {
		fmt.Println(0)
		return
	}
	g := make([][]Node, startIdx[N-1])
	var pushNode func(int, int, int)
	pushNode = func(src, dst, Cap int) {
		g[src] = append(g[src], Node{dst, Cap})
		g[dst] = append(g[dst], Node{src, 0})
	}
	pushNode(0, 1, C[0])
	for i := startIdx[0]; i < startIdx[1]; i++ {
		pushNode(0, i, C[0])
	}
	for i := startIdx[N-2]; i < startIdx[N-1]; i++ {
		pushNode(i, 1, C[N-1])
	}
	for i := 0; i < N-1; i++ {
		if len(train[i]) == 1 {
			pushNode(0, 1, train[i][0].Cap)
		} else {
			pushNode(startIdx[i], 1, train[i][0].Cap)
			for j := 1; j < len(train[i])-1; j++ {
				pushNode(startIdx[i]+j, startIdx[i]+j-1, train[i][j].Cap)
			}
			pushNode(0, startIdx[i+1]-1, train[i][len(train[i])-1].Cap)
		}
	}
	for i := 0; i < N-2; i++ {
		checkUp := 0
		for j := 0; j < len(train[i]); j++ {
			right := train[i][j].T
			for checkUp < len(train[i+1]) {
				if j > 0 && checkUp < len(train[i+1]) {
					if checkUp > 0 {
						pushNode(startIdx[i]+j-1, startIdx[i+1]+checkUp-1, C[i+1])
					} else {
						pushNode(startIdx[i]+j-1, 1, C[i+1])
					}
				}
				if checkUp == len(train[i+1]) {
					break
				}
				if train[i+1][checkUp].S > right {
					break
				}
				checkUp++
				if train[i+1][checkUp-1].S == right {
					break
				}
			}
			if j > 0 && checkUp == len(train[i+1]) && train[i+1][len(train[i+1])-1].S < right {
				pushNode(startIdx[i]+j-1, 0, C[i+1])
			}
		}
		for checkUp < len(train[i+1]) {
			if checkUp > 0 {
				pushNode(0, startIdx[i+1]+checkUp-1, C[i+1])
			}
			checkUp++
		}
	}
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = 1000000007
	}
	visit := make([]bool, len(g))
	qu := &HeapNode{}
	heap.Push(qu, Node{0, 0})
	dist[0] = 0
	for qu.Len() > 0 {
		nd := heap.Pop(qu).(Node)
		if visit[nd.pos] {
			continue
		}
		visit[nd.pos] = true
		if nd.pos == 1 {
			break
		}
		for _, n := range g[nd.pos] {
			npos := n.pos
			ncost := dist[nd.pos] + n.cost
			if ncost < dist[npos] {
				dist[npos] = ncost
				heap.Push(qu, Node{npos, ncost})
			}
		}
	}
	fmt.Println(dist[1])
}

type Node struct {
	pos, cost int
}

type Train struct {
	S, T, Cap int
}

type HeapNode []Node

func (h HeapNode) Len() int            { return len(h) }
func (h HeapNode) Less(i, j int) bool  { return h[i].cost < h[j].cost }
func (h HeapNode) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapNode) Push(x interface{}) { *h = append(*h, x.(Node)) }

func (h *HeapNode) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
