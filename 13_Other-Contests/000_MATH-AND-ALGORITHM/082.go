package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	type pair struct{ x, y int }
	RL := make([]pair, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &RL[i].y, &RL[i].x)
	}
	sort.Slice(RL, func(i, j int) bool {
		return RL[i].x < RL[j].x
	})

	pos := 0
	ans := 0
	for i := 0; i < N; i++ {
		if pos <= RL[i].y {
			ans++
			pos = RL[i].x
		}
	}
	fmt.Println(ans)
}
