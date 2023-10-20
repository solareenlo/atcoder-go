package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [1 << 17]int
	var F [31]int

	var n int
	fmt.Fscan(in, &n)
	res := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		res ^= a[i]
		R := a[i] ^ (a[i] - 1)
		for j := 1; j < 31; j++ {
			if R == (1<<j)-1 {
				F[j-1]++
			}
		}
	}
	cnt := 0
	for i := 29; i >= 0; i-- {
		if res >= (1 << i) {
			if F[i] == 0 {
				fmt.Println(-1)
				return
			}
			res ^= (1 << (i + 1)) - 1
			cnt++
		}
	}
	if res == 0 {
		fmt.Println(cnt)
	} else {
		fmt.Println(-1)
	}
}
