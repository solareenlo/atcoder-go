package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	st := make(map[int]struct{})
	for i := 0; i < N; i++ {
		var id int
		fmt.Fscan(in, &id)
		st[id] = struct{}{}
	}

	ans := make([]int, 0)
	for i := 0; i < M; i++ {
		var id int
		fmt.Fscan(in, &id)
		if _, ok := st[id]; ok {
			continue
		}
		ans = append(ans, id)
	}

	fmt.Println(len(ans))
	for _, e := range ans {
		fmt.Println(e)
	}
}
