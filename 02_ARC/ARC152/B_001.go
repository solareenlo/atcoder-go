package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [200020]int

	var n, l int
	fmt.Fscan(in, &n, &l)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	ans := int(1e18)
	for i, j := 1, n; i <= n; i++ {
		for ; j != 0 && l-a[j] < a[i]; j-- {

		}
		if j != 0 {
			ans = min(ans, 2*l+2*(l-a[j]-a[i]))
		}
		if j < n {
			ans = min(ans, 2*l+2*(a[i]-(l-a[j+1])))
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
