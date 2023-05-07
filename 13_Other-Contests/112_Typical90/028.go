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

	var v [1001][1001]int
	for i := 0; i < N; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		v[a][b]++
		v[a][d]--
		v[c][b]--
		v[c][d]++
	}

	for i := 0; i < 1001; i++ {
		for j := 0; j < 1000; j++ {
			v[i][j+1] += v[i][j]
		}
	}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1001; j++ {
			v[i+1][j] += v[i][j]
		}
	}

	var cnt [200001]int
	for i := 0; i < 1001; i++ {
		for j := 0; j < 1001; j++ {
			cnt[v[i][j]]++
		}
	}

	for i := 0; i < N; i++ {
		fmt.Println(cnt[i+1])
	}
}
