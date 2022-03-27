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
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ok := 1
	if a[1] > 0 {
		ok = 0
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if a[i]-a[i-1] > 1 {
			ok = 0
		} else if a[i] == a[i-1]+1 {
			ans += 1
		} else {
			ans += a[i]
		}
	}

	if ok == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
