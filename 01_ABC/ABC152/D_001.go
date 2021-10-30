package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	ab := [10][10]int{}
	for i := 0; i < n; i++ {
		s := strconv.Itoa(i + 1)
		ab[s[0]-'0'][s[len(s)-1]-'0']++
	}

	res := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			res += ab[i][j] * ab[j][i]
		}
	}

	fmt.Println(res)
}
