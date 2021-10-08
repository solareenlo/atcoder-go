package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x < 4 {
		fmt.Println(1)
	} else {
		maxi := 0
		for i := 2; i < 32; i++ {
			for j := 2; j < 11; j++ {
				if pow(i, j) <= x {
					maxi = max(maxi, pow(i, j))
				}
			}
		}
		fmt.Println(maxi)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
