package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var R, N, M, L int
	fmt.Fscan(in, &R, &N, &M, &L)
	s := make([]int, L)
	for i := 0; i < L; i++ {
		fmt.Fscan(in, &s[i])
	}
	sum, cnt := 0, 0
	for i := 0; i < L; i++ {
		sum += s[i]
		if sum > N {
			fmt.Println("No")
			return
		}
		cnt++
		if cnt == M || sum == N {
			sum = 0
			cnt = 0
			if R == 0 {
				fmt.Println("No")
				return
			}
			R--
		}
	}
	if R > 0 || sum > 0 || cnt > 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
