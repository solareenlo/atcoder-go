package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)
	A := make([][]int, H)
	B := make([][]int, H)
	for i := 0; i < H; i++ {
		A[i] = make([]int, W)
		B[i] = make([]int, W)
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &B[i][j])
		}
	}
	for k := 0; k < 30; k++ {
		uf := NewDsu(2 * (H + W))
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if B[i][j] != -1 {
					if (A[i][j] >> k & 1) == (B[i][j] >> k & 1) {
						uf.Merge(i, H+j)
						uf.Merge(H+W+i, H+W+H+j)
					} else {
						uf.Merge(i, H+W+H+j)
						uf.Merge(H+W+i, H+j)
					}
				}
			}
		}
		for i := 0; i < H+W; i++ {
			if uf.Same(i, H+W+i) {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

type dsu struct {
	n            int
	parentOrSize []int
}

func NewDsu(n int) *dsu {
	d := new(dsu)
	d.n = n
	d.parentOrSize = make([]int, d.n)
	for i := range d.parentOrSize {
		d.parentOrSize[i] = -1
	}
	return d
}

func (d *dsu) Merge(a, b int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	x := d.Leader(a)
	y := d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

func (d *dsu) Same(a, b int) bool {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	return d.Leader(a) == d.Leader(b)
}

func (d *dsu) Leader(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d *dsu) Size(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	return -d.parentOrSize[d.Leader(a)]
}

func (d *dsu) Groups() [][]int {
	leaderBuf := make([]int, d.n)
	groupSize := make([]int, d.n)
	for i := 0; i < d.n; i++ {
		leaderBuf[i] = d.Leader(i)
		groupSize[leaderBuf[i]]++
	}
	result := make([][]int, d.n)
	for i := 0; i < d.n; i++ {
		result[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < d.n; i++ {
		result[leaderBuf[i]] = append(result[leaderBuf[i]], i)
	}
	eraseEmpty := func(a [][]int) [][]int {
		result := make([][]int, 0, len(a))
		for i := range a {
			if len(a[i]) != 0 {
				result = append(result, a[i])
			}
		}
		return result
	}
	return eraseEmpty(result)
}
