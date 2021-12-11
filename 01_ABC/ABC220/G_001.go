package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	x := make([]int, n)
	y := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i], &c[i])
	}

	type pair struct{ x, y int }
	mp := map[pair][]pair{}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			dx := x[i] - x[j]
			dy := y[i] - y[j]
			if dx == 0 {
				dy = 1
			} else {
				if dx < 0 {
					dx = -dx
					dy = -dy
				}
				g := gcd(dx, abs(dy))
				dx /= g
				dy /= g
			}
			mp[pair{dx, dy}] = append(mp[pair{dx, dy}], pair{i, j})
		}
	}

	res := -1
	for a, ps := range mp {
		mp2 := map[int][]pair{}
		for i := range ps {
			mx := x[ps[i].x] + x[ps[i].y]
			my := y[ps[i].x] + y[ps[i].y]
			mp2[mx*a.x+my*a.y] = append(mp2[mx*a.x+my*a.y], ps[i])
		}
		for _, ps2 := range mp2 {
			mp3 := map[pair]int{}
			for i := range ps2 {
				mx := x[ps2[i].x] + x[ps2[i].y]
				my := y[ps2[i].x] + y[ps2[i].y]
				mp3[pair{mx, my}] = max(mp3[pair{mx, my}], c[ps2[i].x]+c[ps2[i].y])
			}
			if len(mp3) < 2 {
				continue
			}
			cs := make([]int, 0)
			for _, nc := range mp3 {
				cs = append(cs, nc)
			}
			sort.Sort(sort.Reverse(sort.IntSlice(cs)))
			res = max(res, cs[0]+cs[1])
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
