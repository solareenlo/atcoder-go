package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var num [200200]map[int]int
	for i := range num {
		num[i] = make(map[int]int)
	}

	var N, M int
	fmt.Fscan(in, &N, &M)
	ans := 0
	for M > 0 {
		M--
		var a, b, l int
		fmt.Fscan(in, &a, &b, &l)
		ans += num[a][2540-l] + num[b][2540-l]
		num[a][l]++
		num[b][l]++
	}
	fmt.Println(ans)
}
