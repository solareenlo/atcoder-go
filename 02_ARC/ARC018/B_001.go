package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				x1 := x[j] - x[i]
				x2 := x[k] - x[i]
				y1 := y[j] - y[i]
				y2 := y[k] - y[i]
				s := x1*y2 - x2*y1
				if s != 0 && s%2 == 0 {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}
