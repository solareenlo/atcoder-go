package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var tweet [1 << 17]int
	M := make([]map[int]int, 1<<17)
	for i := range M {
		M[i] = make(map[int]int)
	}

	var N, MM, K int
	fmt.Fscan(in, &N, &MM, &K)
	cnt := make([]int, N)
	for i := 0; i < MM; i++ {
		var c string
		fmt.Fscan(in, &c)
		var j, k int
		switch c {
		case "t":
			var id int
			fmt.Fscan(in, &id)
			id--
			tweet[id]++
		case "f":
			fmt.Fscan(in, &j, &k)
			j--
			k--
			M[j][k] = tweet[k]
			M[k][j] = tweet[j]
		default:
			fmt.Fscan(in, &j, &k)
			j--
			k--
			cnt[j] += tweet[k] - M[j][k]
			cnt[k] += tweet[j] - M[k][j]
			delete(M[j], k)
			delete(M[k], j)
		}
	}
	for i := 0; i < N; i++ {
		cnt[i] += tweet[i]
		for k, v := range M[i] {
			cnt[i] += tweet[k] - v
		}
	}
	sort.Ints(cnt)
	fmt.Println(cnt[N-K])
}
