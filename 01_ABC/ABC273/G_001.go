package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353
	const N = 5050

	var n int
	fmt.Fscan(in, &n)

	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	var b [N]int
	sum, cnt := 0, 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		sum += b[i]
		if b[i] == 2 {
			cnt++
		}
	}

	var f [N]int
	f[cnt] = 1
	for i := 1; i <= n; i++ {
		sum -= a[i]
		for j := 0; j <= cnt; j++ {
			c1 := sum - 2*j
			if a[i] == 1 {
				f[j] = (f[j]*(c1+1) + f[j+1]*(j+1)) % mod
			}
			if a[i] == 2 {
				f[j] = (f[j]*(c1+1)*(c1+2)/2 + f[j+1]*(j+1)*(c1+1) + f[j+2]*(j+1)*(j+2)/2) % mod
			}
		}
	}

	if sum != 0 {
		fmt.Println(0)
	} else {
		fmt.Println(f[0])
	}
}
