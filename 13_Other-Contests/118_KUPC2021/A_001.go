package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	S := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &S[i])
	}
	sort.Ints(S)
	var T int
	fmt.Fscan(in, &T)
	ans := 1
	last := -1
	for i := 0; i < N; i++ {
		if last >= 0 {
			d1 := last / T
			d2 := S[i] / T
			if d2 > d1 {
				ans++
			}
		}
		last = S[i]
	}
	fmt.Println(ans)
}
