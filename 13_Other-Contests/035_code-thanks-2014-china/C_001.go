package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var N int
	fmt.Fscan(in, &N)
	MA := make(map[pair]int)
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		MA[pair{x[i], y[i]}] = i
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			var k, l int
			ax := x[i]
			ay := y[i]
			bx := x[j]
			by := y[j]
			cx := bx - (by - ay)
			cy := by + (bx - ax)
			if _, ok := MA[pair{cx, cy}]; ok {
				k = MA[pair{cx, cy}]
			} else {
				continue
			}
			dx := ax - (by - ay)
			dy := ay + (bx - ax)
			if _, ok := MA[pair{dx, dy}]; ok {
				l = MA[pair{dx, dy}]
			} else {
				continue
			}
			ans := []int{i, j, k, l}
			sort.Ints(ans)
			fmt.Println(4)
			for _, a := range ans {
				fmt.Println(a + 1)
			}
			return
		}
	}

	fmt.Println(0)
}
