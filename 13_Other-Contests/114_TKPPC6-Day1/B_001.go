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

	mp := make(map[int]struct{})
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		mp[a] = struct{}{}
	}

	fmt.Println(min(N, len(mp)+M))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
