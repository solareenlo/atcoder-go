package main

import "fmt"

func main() {
	var ax, ay, bx, by, n int
	fmt.Scan(&ax, &ay, &bx, &by, &n)

	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		cx := x[i]
		cy := y[i]
		dx := x[(i+1)%n]
		dy := y[(i+1)%n]
		area1 := ax*by + bx*cy + cx*ay - ay*bx - by*cx - cy*ax
		area2 := ax*by + bx*dy + dx*ay - ay*bx - by*dx - dy*ax
		area3 := cx*dy + dx*ay + ax*cy - cy*dx - dy*ax - ay*cx
		area4 := cx*dy + dx*by + bx*cy - cy*dx - dy*bx - by*cx
		if area1*area2 < 0 && area3*area4 < 0 {
			cnt++
		}
	}
	fmt.Println(cnt/2 + 1)
}
