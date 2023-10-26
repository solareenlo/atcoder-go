package main

import "fmt"

func main() {
	var xst, yst, xa, ya, xb, yb, xe, ye int
	fmt.Scan(&xst, &yst, &xa, &ya, &xb, &yb, &xe, &ye)

	for i := 0; i < 33; i++ {
		var tmpXA int
		if xa < 450 {
			tmpXA = 2 + (i * 2)
		} else {
			tmpXA = -2 + (i * -2)
		}
		var tmpYA int
		if ya < 450 {
			tmpYA = i * 1
		} else {
			tmpYA = i * -1
		}
		fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", xa+tmpXA, 500, 60, 60, ya+tmpYA, 500, 60, 60, 2500, 0, 0)
		var tmpXB int
		if xb < 450 {
			tmpXB = i * 1
		} else {
			tmpXB = i * -1
		}
		var tmpYB int
		if yb < 450 {
			tmpYB = 2 + (i * 2)
		} else {
			tmpYB = -2 + (i * -2)
		}
		fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", xb+tmpXB, 500, 60, 60, yb+tmpYB, 500, 60, 60, 0, 2500, 0)
		fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", xe, 500, 60, 60, ye, 500, 60, 60, 0, 0, 5000)
	}
}
