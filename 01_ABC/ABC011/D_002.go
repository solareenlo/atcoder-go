package main

import "fmt"

func main() {
	var n, d, x, y int
	fmt.Scan(&n, &d, &x, &y)

	if x%d != 0 || x%d != 0 {
		fmt.Println(0)
		return
	}

	x, y = x/d, y/d
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	res := 0.0
	for i := x; i <= n-y; i++ {
		RorL := i
		UorD := n - i
		if (RorL-x)%2 != 1 && (UorD-y)%2 != 1 {
			t := float64(1)
			for j := 0; j < RorL; j++ {
				t *= float64(n-j) / float64(RorL-j)
			}
			for j := 0; j < n; j++ {
				t /= 4
			}
			l := (RorL + x) / 2
			for j := 0; j < l; j++ {
				t *= float64(RorL-j) / float64(l-j)
			}
			d := (UorD + y) / 2
			for j := 0; j < d; j++ {
				t *= float64(UorD-j) / float64(d-j)
			}
			res += t
		}
	}
	fmt.Println(res)
}
