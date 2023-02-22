package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type tuple struct {
		x, y, z int
	}

	p := make([][3]int, 300)
	for i := 0; i < 300; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		p[i][0] = x
		p[i][1] = y
		p[i][2] = i + 1
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i][0] == p[j][0] {
			if p[i][1] == p[j][1] {
				return p[i][2] < p[j][2]
			}
			return p[i][1] < p[j][1]
		}
		return p[i][0] < p[j][0]
	})
	fmt.Println(100)
	for i := 0; i < 100; i++ {
		fmt.Printf("%d %d %d\n", p[i*3][2], p[i*3+1][2], p[i*3+2][2])
	}
}
