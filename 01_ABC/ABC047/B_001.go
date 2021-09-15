package main

import "fmt"

func main() {
	var w, h, n, x, y, a int
	fmt.Scan(&w, &h, &n)

	minX, maxX, minY, maxY := 0, w, 0, h
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y, &a)
		switch a {
		case 1:
			minX = max(minX, x)
		case 2:
			maxX = min(maxX, x)
		case 3:
			minY = max(minY, y)
		case 4:
			maxY = min(maxY, y)
		}
	}

	fmt.Println(max(0, maxX-minX) * max(0, maxY-minY))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
