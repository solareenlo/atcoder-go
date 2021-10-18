package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func scanString() string { sc.Scan(); return sc.Text() }
func scanInt() int       { a, _ := strconv.Atoi(scanString()); return a }

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), 1048576)
	defer out.Flush()

	n, q := scanInt(), scanInt()

	I := make([][]int, 2*n)
	for i := 0; i < n; i++ {
		s, t, x := scanInt(), scanInt(), scanInt()
		I[i] = make([]int, 4)
		I[i][0], I[i][1], I[i][2], I[i][3] = s-x, 1, i, x
		I[i+n] = make([]int, 4)
		I[i+n][0], I[i+n][1], I[i+n][2], I[i+n][3] = t-x, 0, i, x
	}
	sort.Slice(I, func(i, j int) bool {
		return I[i][0] < I[j][0]
	})

	id := 0
	pq := &Heap{}
	m := make(map[int]int)
	for i := 0; i < q; i++ {
		d := scanInt()
		for id < 2*n && I[id][0] <= d {
			if I[id][1] == 1 {
				heap.Push(pq, I[id][3])
			} else {
				m[I[id][3]]++
			}
			id++
		}
		for pq.Len() > 0 && m[(*pq)[0]] > 0 {
			m[heap.Pop(pq).(int)]--
		}
		if pq.Len() == 0 {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, (*pq)[0])
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
