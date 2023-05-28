package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func main() {
	in := bufio.NewReader(os.Stdin)

	var k int
	fmt.Fscan(in, &n, &k)

	ans := 0
	var f [200005]int
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		s := sum(n - x + 1)
		Add(n - x + 1)
		if s == 0 {
			continue
		}
		s = (s - 1) / (k - 1)
		ans += s
		if i-(k-1)*s > k {
			ans++
		} else {
			f[s] = 1
		}
	}
	for i := 0; i < n; i++ {
		ans += f[i]
	}
	fmt.Println(ans)
}

var t [200005]int

func sum(x int) int {
	res := 0
	for ; x > 0; x -= x & -x {
		res += t[x]
	}
	return res
}

func Add(x int) {
	for ; x <= n; x += x & -x {
		t[x]++
	}
}
