package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 300000

	var n int
	fmt.Fscan(in, &n)

	as := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &as[i])
	}

	ais := make([]pair, n)
	for i := 0; i < n; i++ {
		ais[i] = pair{as[i], i}
	}
	sortPair(ais)

	bs := make([]int, N)
	for i := 0; i < n; i++ {
		bs[ais[i].y] = i
	}

	for i := 0; i < n; i++ {
		if i+1 < n {
			fmt.Printf("%d ", bs[i]+1)
		} else {
			fmt.Println(bs[i] + 1)
		}
	}
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
