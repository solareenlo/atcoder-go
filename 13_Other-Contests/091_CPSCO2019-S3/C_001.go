package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var v [1000001]int
	for i := 0; i < n; i++ {
		var s, t int
		fmt.Fscan(in, &s, &t)
		v[s]++
		v[t]--
	}
	ans := 0
	for i := 1; i <= 1000000; i++ {
		v[i] += v[i-1]
		if v[i-1] == 0 && v[i] != 0 {
			ans++
		}
	}
	fmt.Println(ans)
}
