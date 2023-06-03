package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	var a [30]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	ans := false
	fr := make(map[int]int)
	for i := 0; i < (1 << 15); i++ {
		sum := 0
		for j := 0; j < 15; j++ {
			if (i & (1 << j)) != 0 {
				sum += a[j]
			}
		}
		fr[sum] = 1
	}
	for i := 0; i < (1 << 15); i++ {
		sum := 0
		for j := 15; j < 30; j++ {
			if (i & (1 << (j - 15))) != 0 {
				sum += a[j]
			}
		}
		if fr[k-sum] == 1 {
			ans = true
		}
	}
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
