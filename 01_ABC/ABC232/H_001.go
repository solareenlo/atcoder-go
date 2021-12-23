package main

import "fmt"

type P struct{ x, y int }

var (
	res        = make([]P, 0)
	h, w, a, b int
)

func solve(h, w, a, b int) []P {
	ret := make([]P, 0)
	if h == 2 {
		for i := 1; i <= b-1; i++ {
			ret = append(ret, P{1, i})
			ret = append(ret, P{2, i})
		}
		ret = append(ret, P{3 - a, b})
		for i := b + 1; i <= w; i++ {
			ret = append(ret, P{1, i})
		}
		for i := w; i >= b+1; i-- {
			ret = append(ret, P{2, i})
		}
		ret = append(ret, P{a, b})
	} else if h > 2 && w == 2 || b == 1 || a == h && b == 2 {
		ret = solve(w, h, b, a)
		for i := range ret {
			ret[i].x, ret[i].y = ret[i].y, ret[i].x
		}
	} else {
		for i := 1; i <= h; i++ {
			ret = append(ret, P{i, 1})
		}
		res := solve(h, w-1, h-a+1, b-1)
		for i := range res {
			res[i].x = h + 1 - res[i].x
			res[i].y++
		}
		ret = append(ret, res...)
	}
	return ret
}

func main() {
	fmt.Scan(&h, &w, &a, &b)
	ans := solve(h, w, a, b)
	for _, p := range ans {
		fmt.Println(p.x, p.y)
	}
}
