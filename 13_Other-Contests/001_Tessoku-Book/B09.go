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

	var CS, CSX [1502][1502]int
	for N > 0 {
		N--
		var A, B, C, D int
		fmt.Fscan(in, &A, &B, &C, &D)
		A++
		B++
		C++
		D++
		CS[A][B]++
		CS[A][D]--
		CS[C][B]--
		CS[C][D]++
	}

	cnt := 0
	for i := 1; i <= 1501; i++ {
		for j := 1; j <= 1501; j++ {
			CSX[i][j] = CS[i][j] + CSX[i-1][j] + CSX[i][j-1] - CSX[i-1][j-1]
			if CSX[i][j] != 0 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
