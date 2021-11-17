package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	y *= 2

	res := 1 << 60
	cnt := 0
	for {
		res = min(res, cnt+abs(y-x*2)/2)
		tmp := 0
		if y%4 == 3 {
			tmp = 1
		}
		y = y/2 + tmp
		cnt += 1 + y%2
		if y < x {
			break
		}
	}

	fmt.Println(res)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
