package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// 元論文: <http://www.ics.uci.edu/~eppstein/pubs/Epp-SJC-98.pdf>

const InfWeight = 0x3f3f3f3f3f3f3f3f

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	edges := make([]Edge, 0)
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(in, &u, &v, &c)
		edges = append(edges, Edge{u, To{v, c}})
	}
	g := NewWeightedStaticGraph(n, edges)
	kShortestPaths := NewKShortestPaths(*g, 0)
	lengths := kShortestPaths.getDistances(0, k+1)
	for i := 1; i <= k; i++ {
		if len(lengths) <= i {
			fmt.Println(-1)
		} else {
			fmt.Println(lengths[i])
		}
	}
}

type Graph struct {
	n       int
	tos     []To
	offsets []int
}

func NewWeightedStaticGraph(n int, edges []Edge) *Graph {
	g := new(Graph)
	g.n = n
	m := len(edges)
	g.tos = make([]To, m+1)
	g.offsets = make([]int, n+1)
	for e := 0; e < m; e++ {
		g.offsets[edges[e].x]++
	}
	for v := 1; v <= n; v++ {
		g.offsets[v] += g.offsets[v-1]
	}
	for e := 0; e < m; e++ {
		g.offsets[edges[e].x]--
		g.tos[g.offsets[edges[e].x]] = edges[e].to
	}
	return g
}

func (g Graph) numVertices() int { return g.n }

type Edge struct {
	x  int
	to To
}

type To struct {
	to, w int
}

type KShortestPaths struct {
	hBestNodes            []HBestNode
	hBestRoots            []int
	hasSidetrack          []bool
	bestSidetracks        []To
	sidetrackHeaps        []To
	sidetrackHeapOffsets  []int
	toDestinationDistance []int
	vertexIndices         []int
}

func NewKShortestPaths(g Graph, destination int) *KShortestPaths {
	K := new(KShortestPaths)
	n := g.numVertices()
	dist := make([]int, n)
	treeParent := make([]int, n)
	treeOrder := make([]int, 0)
	K.buildShortestPathTree(g, destination, &dist, &treeParent, &treeOrder)
	K.buildMain(g, dist, treeParent, treeOrder)
	return K
}

func (K *KShortestPaths) buildMain(g Graph, dist, treeParent, treeOrder []int) {
	n := g.numVertices()
	K.hBestNodes = make([]HBestNode, 0)
	K.hBestRoots = make([]int, n)
	for i := range K.hBestRoots {
		K.hBestRoots[i] = -1
	}
	K.hasSidetrack = make([]bool, n)
	K.bestSidetracks = make([]To, n)
	K.sidetrackHeaps = make([]To, 0)
	K.sidetrackHeapOffsets = make([]int, len(treeOrder)+1)
	K.toDestinationDistance = make([]int, n)
	for i := 0; i < n; i++ {
		K.toDestinationDistance[i] = dist[i]
	}
	K.vertexIndices = make([]int, n)
	for i := range K.vertexIndices {
		K.vertexIndices[i] = -1
	}
	hBestSizes := make([]int, n)

	for idx := 0; idx < len(treeOrder); idx++ {
		v := treeOrder[idx]
		p := treeParent[v]
		K.vertexIndices[v] = idx

		var bestSidetrack To
		bestSidetrack.w = InfWeight
		treeEdgeIndex := -1
		bestEdgeIndex := -1
		sawTreeEdge := false
		ei := 0
		for i := g.offsets[v]; i < g.offsets[v+1]; i++ {
			u := g.tos[i].to
			delta := g.tos[i].w - (dist[v] - dist[u])
			if !sawTreeEdge && treeParent[v] == u && delta == 0 {
				treeEdgeIndex = ei
				sawTreeEdge = true
			} else if bestSidetrack.w > delta {
				K.hasSidetrack[v] = true
				bestEdgeIndex = ei
				bestSidetrack = To{u, delta}
			}
			ei++
		}
		K.bestSidetracks[v] = bestSidetrack
		if p == -1 {
			K.hBestRoots[v] = -1
			hBestSizes[v] = 0
		} else {
			K.hBestRoots[v] = K.hBestRoots[p]
			hBestSizes[v] = hBestSizes[p]
		}
		if K.hasSidetrack[v] {
			K.hBestRoots[v] = K.insertToBestHeap(K.hBestRoots[v], hBestSizes[v], v)
			hBestSizes[v]++
		}
		K.sidetrackHeapOffsets[idx] = len(K.sidetrackHeaps)
		ei = 0
		for i := g.offsets[v]; i < g.offsets[v+1]; i++ {
			u := g.tos[i].to
			delta := g.tos[i].w - (dist[v] - dist[u])
			if ei != treeEdgeIndex && ei != bestEdgeIndex && delta < InfWeight {
				K.sidetrackHeaps = append(K.sidetrackHeaps, To{u, delta})
			}
			ei++
		}
		tmp := K.sidetrackHeaps[K.sidetrackHeapOffsets[idx]:]
		tmp = makeHeap(tmp)
		for i := range tmp {
			K.sidetrackHeaps[i+K.sidetrackHeapOffsets[idx]] = tmp[i]
		}
	}
	K.sidetrackHeapOffsets[len(treeOrder)] = len(K.sidetrackHeaps)
}

func makeHeap(v []To) []To {
	q := &HeapTo{}
	for i := range v {
		heap.Push(q, v[i])
	}
	return (*q)
}

type HeapTo []To

func (h HeapTo) Len() int            { return len(h) }
func (h HeapTo) Less(i, j int) bool  { return h[i].w < h[j].w }
func (h HeapTo) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTo) Push(x interface{}) { *h = append(*h, x.(To)) }

func (h *HeapTo) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (K KShortestPaths) getDistances(source, k int) []int {
	lengths := make([]int, 0)
	shortestDistance := K.toDestinationDistance[source]
	if shortestDistance == InfWeight || k <= 0 {
		return lengths
	}
	lengths = append(lengths, shortestDistance)
	hBests := len(K.hBestNodes)
	q := &HeapState{}
	root := K.hBestRoots[source]
	if root != -1 {
		rootint := K.hBestNodes[root].vertex
		heap.Push(q, SearchState{K.bestSidetracks[rootint].w, rootint, root})
	}
	for q.Len() > 0 && len(lengths) < k {
		tmp := heap.Pop(q).(SearchState)
		d := tmp.dist
		v := tmp.vertex
		vIndex := K.vertexIndices[v]
		index := tmp.index
		lengths = append(lengths, shortestDistance+d)
		if index < hBests {
			node := K.hBestNodes[index]
			edge := K.bestSidetracks[v]
			e := d - edge.w
			// H_bestの子たち
			left := node.left
			right := node.right
			if left != -1 {
				leftint := K.hBestNodes[left].vertex
				heap.Push(q, SearchState{e + K.bestSidetracks[leftint].w, leftint, left})
			}
			if right != -1 {
				rightint := K.hBestNodes[right].vertex
				heap.Push(q, SearchState{e + K.bestSidetracks[rightint].w, rightint, right})
			}
			// 辺の行き先
			u := edge.to
			uRoot := K.hBestRoots[u]
			if uRoot != -1 {
				rootint := K.hBestNodes[uRoot].vertex
				heap.Push(q, SearchState{d + K.bestSidetracks[rootint].w, rootint, uRoot})
			}
			// best以外の辺
			heapRoot := K.sidetrackHeapOffsets[vIndex]
			if heapRoot != K.sidetrackHeapOffsets[vIndex+1] {
				heap.Push(q, SearchState{e + K.sidetrackHeaps[heapRoot].w, v, hBests + heapRoot})
			}
		} else {
			edge := K.sidetrackHeaps[index-hBests]
			e := d - edge.w
			heapBase := K.sidetrackHeapOffsets[vIndex]
			heapEnd := K.sidetrackHeapOffsets[vIndex+1]
			localIndex := index - hBests - heapBase
			left := heapBase + localIndex*2 + 1
			right := heapBase + localIndex*2 + 2
			// H_outの子たち
			if left < heapEnd {
				heap.Push(q, SearchState{e + K.sidetrackHeaps[left].w, v, hBests + left})
			}
			if right < heapEnd {
				heap.Push(q, SearchState{e + K.sidetrackHeaps[right].w, v, hBests + right})
			}
			// 辺の行き先
			u := edge.to
			uRoot := K.hBestRoots[u]
			if uRoot != -1 {
				rootint := K.hBestNodes[uRoot].vertex
				heap.Push(q, SearchState{d + K.bestSidetracks[rootint].w, rootint, uRoot})
			}
		}
	}
	return lengths
}

type HeapState []SearchState

func (h HeapState) Len() int            { return len(h) }
func (h HeapState) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h HeapState) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapState) Push(x interface{}) { *h = append(*h, x.(SearchState)) }

func (h *HeapState) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (K KShortestPaths) buildShortestPathTree(g Graph, t int, dist, prev, order *[]int) {
	n := g.numVertices()
	h := transposeGraph(g)
	for v := 0; v < n; v++ {
		(*dist)[v] = InfWeight
		(*prev)[v] = -1
	}
	(*dist)[t] = 0
	*order = make([]int, 0)
	visited := make([]bool, n)
	q := &HeapDistPair{}
	heap.Push(q, DistPair{0, t})
	for q.Len() > 0 {
		tmp := heap.Pop(q).(DistPair)
		v := tmp.y
		if visited[v] {
			continue
		}
		visited[v] = true
		*order = append(*order, v)
		for i := h.offsets[v]; i < h.offsets[v+1]; i++ {
			d := (*dist)[v] + h.tos[i].w
			u := h.tos[i].to
			if d < (*dist)[u] {
				(*prev)[u] = v
				(*dist)[u] = d
				heap.Push(q, DistPair{d, u})
			}
		}
	}
}

func transposeGraph(g Graph) *Graph {
	n := g.numVertices()
	edges := make([]Edge, 0)
	for v := 0; v < n; v++ {
		for i := g.offsets[v]; i < g.offsets[v+1]; i++ {
			edges = append(edges, Edge{g.tos[i].to, To{v, g.tos[i].w}})
		}
	}
	return NewWeightedStaticGraph(n, edges)
}

type DistPair struct {
	x, y int
}

type HeapDistPair []DistPair

func (h HeapDistPair) Len() int            { return len(h) }
func (h HeapDistPair) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapDistPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapDistPair) Push(x interface{}) { *h = append(*h, x.(DistPair)) }

func (h *HeapDistPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (K *KShortestPaths) insertToBestHeap(rootNode, heapSize, v int) int {
	if rootNode == -1 {
		return K.newHBestNode(v, -1, -1)
	}
	newPosition := heapSize + 1
	height := 0
	for (1 << height) <= newPosition {
		height++
	}
	levelNodes := make([]int, 32)
	currentNode := rootNode
	insertHeight := 0
	for h := height - 1; h > 0; h-- {
		levelNodes[h] = currentNode
		cnt := K.hBestNodes[currentNode]
		if insertHeight == 0 && K.bestSidetracks[v].w <= K.bestSidetracks[cnt.vertex].w {
			insertHeight = h
		}
		if ((newPosition >> (h - 1)) & 1) != 0 {
			currentNode = max(0, cnt.right)
		} else {
			currentNode = max(0, cnt.left)
		}
	}
	currentNode = -1
	for h := 0; h <= insertHeight; h++ {
		var u int
		if h == insertHeight {
			u = v
		} else {
			u = K.hBestNodes[levelNodes[h+1]].vertex
		}
		if h == 0 {
			currentNode = K.newHBestNode(u, -1, -1)
		} else if ((newPosition >> (h - 1)) & 1) != 0 {
			currentNode = K.newHBestNode(u, K.hBestNodes[levelNodes[h]].left, currentNode)
		} else {
			currentNode = K.newHBestNode(u, currentNode, K.hBestNodes[levelNodes[h]].right)
		}
	}
	for h := insertHeight + 1; h < height; h++ {
		node := K.hBestNodes[levelNodes[h]]
		if ((newPosition >> (h - 1)) & 1) != 0 {
			currentNode = K.newHBestNode(node.vertex, node.left, currentNode)
		} else {
			currentNode = K.newHBestNode(node.vertex, currentNode, node.right)
		}
	}
	return currentNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (K *KShortestPaths) newHBestNode(v, left, right int) int {
	newNode := HBestNode{v, left, right}
	K.hBestNodes = append(K.hBestNodes, newNode)
	return len(K.hBestNodes) - 1
}

type HBestNode struct { // H_bestを構成するノード(永続的2分ヒープ)
	vertex      int // どの頂点のbestか
	left, right int // ヒープの子。ポインタの代わりにインデックス
}

type SearchState struct {
	dist   int
	vertex int
	index  int
}
