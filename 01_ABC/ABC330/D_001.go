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
	var grid [2020]string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &grid[i])
	}

	var row, col [2020]int
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if grid[y][x] == 'o' {
				row[y] += 1
				col[x] += 1
			}
		}
	}

	ans := 0
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if grid[y][x] == 'o' {
				ans += (row[y] - 1) * (col[x] - 1)
			}
		}
	}
	fmt.Println(ans)
}
