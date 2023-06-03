package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	lr := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &lr[i].y, &lr[i].x)
	}
	sortPair(lr)

	ans := 0
	now := 0
	for i := 0; i < n; i++ {
		if now <= lr[i].y {
			ans++
			now = lr[i].x
		}
	}
	fmt.Println(ans)
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
