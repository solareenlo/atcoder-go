package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var A [1501][1501]int
	for n > 0 {
		n--
		var x, y int
		fmt.Fscan(in, &y, &x)
		A[y][x]++
	}
	for y := 1; y < 1501; y++ {
		for x := 1; x < 1501; x++ {
			A[y][x] += A[y-1][x] + A[y][x-1] - A[y-1][x-1]
		}
	}
	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var y, x, y2, x2 int
		fmt.Fscan(in, &y, &x, &y2, &x2)
		x--
		y--
		fmt.Println(A[y2][x2] + A[y][x] - A[y2][x] - A[y][x2])
	}
}
