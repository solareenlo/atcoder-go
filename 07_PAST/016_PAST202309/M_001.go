package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type point struct {
		x, y int
	}

	p := make([]point, 4)
	for i := 0; i < 4; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
	}
	ok := true
	for j := 0; j < 2; j++ {
		for i := 0; i < 2; i++ {
			p[i], p[i+2] = p[i+2], p[i]
		}
		A := p[0].y - p[1].y
		B := p[1].x - p[0].x
		C := -A*p[0].x - B*p[0].y
		if 0 < (A*p[2].x+B*p[2].y+C)*(A*p[3].x+B*p[3].y+C) {
			ok = false
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			p[j].x, p[j].y = p[j].y, p[j].x
		}
		for k := 0; k < 2; k++ {
			for l := 0; l < 2; l++ {
				p[l], p[l+2] = p[l+2], p[l]
			}
			A := max(p[0].x, p[1].x)
			B := min(p[2].x, p[3].x)
			if A < B {
				ok = false
			}
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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
