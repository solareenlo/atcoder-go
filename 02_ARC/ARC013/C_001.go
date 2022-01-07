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

	g := 0
	for i := 0; i < n; i++ {
		var X, Y, Z, m int
		fmt.Fscan(in, &X, &Y, &Z, &m)
		minX, maxX := X, 0
		minY, maxY := Y, 0
		minZ, maxZ := Z, 0
		for j := 0; j < m; j++ {
			var x, y, z int
			fmt.Fscan(in, &x, &y, &z)
			minX = min(minX, x)
			maxX = max(maxX, x)
			minY = min(minY, y)
			maxY = max(maxY, y)
			minZ = min(minZ, z)
			maxZ = max(maxZ, z)
		}
		g ^= minX
		g ^= X - maxX - 1
		g ^= minY
		g ^= Y - maxY - 1
		g ^= minZ
		g ^= Z - maxZ - 1
	}

	if g != 0 {
		fmt.Println("WIN")
	} else {
		fmt.Println("LOSE")
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
