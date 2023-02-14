package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX = 1e7

func main() {
	in := bufio.NewReader(os.Stdin)

	var dx = [9]float64{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	var dy = [9]float64{-1, 0, 1, -1, 0, 1, -1, 0, 1}

	var n int
	var r float64
	fmt.Fscan(in, &n, &r)
	ans := 0
	p := make(map[int]bool)
	for i := 1; i <= n; i++ {
		var x, y float64
		fmt.Fscan(in, &x, &y)
		x += MAX
		y += MAX
		if _, ok := p[dis(int(x/r), int(y/r))]; !ok {
			ans++
		}
		for j := 0; j < 9; j++ {
			p[dis(int((x/r)+dx[j]), int((y/r)+dy[j]))] = true
		}
	}
	fmt.Println(ans)
}

func dis(x, y int) int { return x*MAX + y }
