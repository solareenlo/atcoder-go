package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}
	ans := -1
	for i := 1; i < n; i++ {
		if v[i]-v[i-1] <= d {
			ans = v[i]
			break
		}
	}
	fmt.Println(ans)
}
