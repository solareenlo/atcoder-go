package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 52 * 52 * 52

func char_to_int(c byte) int {
	if 'A' <= c && c <= 'Z' {
		return int(c - 'A')
	} else {
		return int(c-'a') + 26
	}
}

func string_to_int(a, b, c byte) int {
	return char_to_int(a)*52*52 + char_to_int(b)*52 + char_to_int(c)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	type pair struct{ x, y int }
	edge := make([]pair, n)
	revGraph := make([][]int, M)
	cnt := make([]int, M)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		edge[i] = pair{string_to_int(s[0], s[1], s[2]), string_to_int(s[len(s)-3], s[len(s)-2], s[len(s)-1])}
		cnt[edge[i].x]++
		revGraph[edge[i].y] = append(revGraph[edge[i].y], edge[i].x)
	}

	res := make([]int, M)
	for i := range res {
		res[i] = -1
	}
	que := make([]int, 0)
	for i := 0; i < M; i++ {
		if cnt[i] == 0 {
			res[i] = 0
			que = append(que, i)
		}
	}
	for len(que) > 0 {
		t := que[0]
		que = que[1:]
		for _, x := range revGraph[t] {
			if res[x] == -1 {
				cnt[x]--
				if res[t] == 0 {
					res[x] = 1
					que = append(que, x)
				} else if cnt[x] == 0 {
					res[x] = 0
					que = append(que, x)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		switch res[edge[i].y] {
		case -1:
			fmt.Fprintln(out, "Draw")
		case 0:
			fmt.Fprintln(out, "Takahashi")
		case 1:
			fmt.Fprintln(out, "Aoki")
		}
	}
}
