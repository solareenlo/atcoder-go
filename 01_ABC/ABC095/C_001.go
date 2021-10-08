package main

import "fmt"

func main() {
	var a, b, c, x, y int
	fmt.Scan(&a, &b, &c, &x, &y)

	res := 0
	if a+b > 2*c {
		if x == y {
			res = 2 * c * x
		} else if x > y {
			res += 2 * c * y
			if a >= 2*c {
				res += 2 * c * (x - y)
			} else {
				res += a * (x - y)
			}
		} else if x < y {
			res += 2 * c * x
			if b >= 2*c {
				res += 2 * c * (y - x)
			} else {
				res += b * (y - x)
			}
		}
	} else if a+b <= 2*c {
		res += a*x + b*y
	}
	fmt.Println(res)
}
