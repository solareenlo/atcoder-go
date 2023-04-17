package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, p, q int
	fmt.Fscan(in, &n, &p, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	if (p+q)%2 != 0 {
		fmt.Println(0)
		return
	}

	a1 := (p + q) / 2
	ans := 0
	cnt := 0
	f := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		if a[i] == a1 {
			ans += cnt
		}
		cnt += f[(p-a1)-a[i]]
		f[a[i]]++
	}
	fmt.Println(ans)
}
