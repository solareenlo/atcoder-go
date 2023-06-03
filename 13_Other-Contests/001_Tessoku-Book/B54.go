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
	T := make(map[int]int)
	ans := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		ans += T[x]
		T[x] = T[x] + 1
	}
	fmt.Println(ans)
}
