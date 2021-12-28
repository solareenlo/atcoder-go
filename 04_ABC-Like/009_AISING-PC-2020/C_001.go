package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		cnt := 0
		m := int(math.Sqrt(float64(n)))
		for x := 1; x <= m; x++ {
			for y := 1; y <= m; y++ {
				d := math.Sqrt(float64(-2*x*y - 3*pow(x, 2) - 3*pow(y, 2) + 4*i))
				if math.Floor(d) == d {
					z := (float64(-x-y) + d) / 2
					if math.Floor(z) == z && z > 0 {
						if pow(x, 2)+pow(y, 2)+pow(int(z), 2)+x*y+y*int(z)+int(z)*x == i {
							cnt++
						}
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}
