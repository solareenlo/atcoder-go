package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	var s [9]string
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &s[i])
	}
	x := 1
	y := 1
	px := -1
	py := -1
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	for i := 0; i < 1000; i++ {
		if abs(x-px)+abs(y-py) > 1 {
			fmt.Println("-")
		} else {
			for r := 0; r < 4; r++ {
				tx := x + dx[r]
				ty := y + dy[r]
				if s[tx][ty] == '#' || px == tx && py == ty {
					continue
				}
				x = tx
				y = ty
				fmt.Println(string("RDLU"[r]))
				break
			}
		}
		fmt.Scan(&px, &py)
		px--
		py--
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
