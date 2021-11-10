package main

import "fmt"

func main() {
	var x, k, d int
	fmt.Scan(&x, &k, &d)

	x = abs(x)
	div := x / d

	res := 0
	if x == d {
		if k%2 != 0 {
			res = 0
		} else {
			res = d
		}
	} else if x < d {
		if k%2 != 0 {
			res = abs(x - d)
		} else {
			res = x
		}
	} else if div > k {
		res = abs(x - d*k)
	} else {
		if (k-div)%2 != 0 {
			res = abs(x - d*(div+1))
		} else {
			res = abs(x - d*div)
		}
	}

	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
