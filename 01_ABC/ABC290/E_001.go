package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var f [200200]int

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	ans := 0
	for i := 1; i <= n; i++ {
		ans += (n - i + 1) * (i / 2)
		fmt.Fscan(in, &a[i])
	}
	if (n & 1) != 0 {
		f[a[n/2+1]]++
	}
	for i, j := n/2, n/2+1+n%2; i > 0; i, j = i-1, j+1 {
		tmp := 0
		if a[i] == a[j] {
			tmp = 1
		}
		ans -= (f[a[i]] + f[a[j]] + tmp) * i
		f[a[i]]++
		f[a[j]]++
	}
	fmt.Println(ans)
}
