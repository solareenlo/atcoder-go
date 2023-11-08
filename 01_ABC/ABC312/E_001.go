package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var id [111][111][111]int
	dx := [3]int{0, 0, 1}
	dy := [3]int{1, 0, 0}
	dz := [3]int{0, 1, 0}

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		var a, b, c, d, e, f int
		fmt.Fscan(in, &a, &b, &c, &d, &e, &f)
		for x := a; x < d; x++ {
			for y := b; y < e; y++ {
				for z := c; z < f; z++ {
					id[x][y][z] = i
				}
			}
		}
	}
	ans := make([]map[int]bool, 200200)
	for i := range ans {
		ans[i] = make(map[int]bool)
	}
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			for z := 0; z < 100; z++ {
				if id[x][y][z] != 0 {
					for i := 0; i < 3; i++ {
						if id[x+dx[i]][y+dy[i]][z+dz[i]] != 0 && id[x][y][z]^id[x+dx[i]][y+dy[i]][z+dz[i]] != 0 {
							ans[id[x][y][z]][id[x+dx[i]][y+dy[i]][z+dz[i]]] = true
							ans[id[x+dx[i]][y+dy[i]][z+dz[i]]][id[x][y][z]] = true
						}
					}
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Println(len(ans[i]))
	}
}
