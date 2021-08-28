package main

import "fmt"

func main() {
	var n, D, X, Y int
	fmt.Scan(&n, &D, &X, &Y)

	if X%D != 0 || Y%D != 0 {
		fmt.Println(0)
		return
	}

	x, y := float64(X/D), float64(Y/D)
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	d := float64(float64(n)-x-y) / 2
	if d < 0 || ((n-X/D-Y/D)%2) != 0 {
		fmt.Println(0)
		return
	}

	res := 1.0
	for i := 0; i < n; i++ {
		res /= 4.0
		if d-float64(i) > 0 {
			res *= float64((n-i)*(n-i)) / ((d - float64(i)) * (y + d - float64(i)))
		} else if y+d-float64(i) > 0 {
			res *= float64(n-i) / (y + d - float64(i))
		}
	}
	fmt.Println(res)
}
