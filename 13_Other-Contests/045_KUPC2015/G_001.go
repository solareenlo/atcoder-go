package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type BIT struct {
	size_ int
	bit   []int
}

func (b *BIT) init(sz int) {
	b.size_ = sz + 2
	b.bit = make([]int, b.size_+2)
}

func (b *BIT) add(pos, x int) {
	pos++
	for pos <= b.size_ {
		b.bit[pos] += x
		pos += (pos & -pos)
	}
}

func (b *BIT) sum(pos int) int {
	pos++
	s := 0
	for pos >= 1 {
		s += b.bit[pos]
		pos -= (pos & -pos)
	}
	return s
}

type PQ []int

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i] > pq[j] }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	var A, B [1 << 18]int
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &B[i])
	}

	var X BIT
	X.init(N + 2)
	var Z = make([][2]int, 0, N)
	for i := 1; i <= N; i++ {
		Z = append(Z, [2]int{A[i], i})
	}
	sort.Slice(Z, func(i, j int) bool { return Z[i][0] < Z[j][0] })

	cx := len(Z) - 1
	ans := int64(0)
	for i := 1; i <= N; i++ {
		X.add(i, 1)
	}
	var Q PQ
	heap.Init(&Q)
	for i := N; i >= 1; i-- {
		for cx >= 0 && Z[cx][0] >= B[i] {
			heap.Push(&Q, Z[cx][1])
			cx--
		}
		if Q.Len() == 0 {
			fmt.Println("-1")
			return
		}
		r := heap.Pop(&Q).(int)
		s := X.sum(N) - X.sum(r)
		ans += int64(s)
		X.add(r, -1)
	}
	fmt.Println(ans)
}
