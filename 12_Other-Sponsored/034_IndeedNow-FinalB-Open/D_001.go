package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)
	var sx, sy int
	fmt.Fscan(in, &sx, &sy)
	var gx, gy int
	fmt.Fscan(in, &gx, &gy)
	sx--
	sy--
	gx--
	gy--
	P := make([][]int, H)
	for i := range P {
		P[i] = make([]int, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &P[i][j])
		}
	}

	var flatten func(int, int) int
	flatten = func(h, w int) int {
		return h*W + w
	}

	Bonus := make([]tuple, 0)
	for h := 0; h < H; h++ {
		for w := 0; w < W-1; w++ {
			Bonus = append(Bonus, tuple{P[h][w] * P[h][w+1], flatten(h, w), flatten(h, w+1)})
		}
	}
	for w := 0; w < W; w++ {
		for h := 0; h < H-1; h++ {
			Bonus = append(Bonus, tuple{P[h][w] * P[h+1][w], flatten(h, w), flatten(h+1, w)})
		}
	}
	sortTuple(Bonus)
	uf := NewDsu(H * W)
	ans := 0
	for _, tmp := range Bonus {
		val := tmp.x
		p := tmp.y
		q := tmp.z
		if !uf.Same(p, q) {
			uf.Merge(p, q)
			ans += val
		}
	}
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			ans += P[h][w]
		}
	}
	fmt.Println(ans)
}

type tuple struct {
	x, y, z int
}

func sortTuple(tup []tuple) {
	sort.Slice(tup, func(i, j int) bool {
		if tup[i].x == tup[j].x {
			if tup[i].y == tup[j].y {
				return tup[i].z > tup[j].z
			}
			return tup[i].y > tup[j].y
		}
		return tup[i].x > tup[j].x
	})
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
