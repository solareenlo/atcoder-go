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

	a := make([]int, n+1)
	b := make([]int, 100005)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		b[a[i]]++
	}

	if a[1] != 0 || b[0] != 1 {
		fmt.Println(0)
		return
	}

	ans := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= b[i]; j++ {
			ans = ans * b[i-1] % 998244353
		}
	}

	fmt.Println(ans)
}
